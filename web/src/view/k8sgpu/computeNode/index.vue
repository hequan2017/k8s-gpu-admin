<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDrawer">新增节点</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column align="left" label="ID" prop="ID" width="60" />
        <el-table-column align="left" label="节点名称" prop="name" width="120" />
        <el-table-column align="left" label="区域" prop="region" width="100" />
        <el-table-column align="left" label="CPU" prop="cpu" width="80" />
        <el-table-column align="left" label="内存(GB)" prop="memory" width="100" />
        <el-table-column align="left" label="显卡" prop="gpuName" width="120" />
        <el-table-column align="left" label="显卡数量" prop="gpuCount" width="100" />
        <el-table-column align="left" label="公网IP" prop="publicIP" width="140" />
        <el-table-column align="left" label="内网IP" prop="privateIP" width="140" />
        <el-table-column align="left" label="是否上架" width="100">
          <template #default="scope">
            <el-switch
              v-model="scope.row.isPublished"
              @change="togglePublished(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" width="180">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              @click="updateComputeNode(scope.row)"
            >编辑</el-button>
            <el-button
              type="danger"
              link
              icon="delete"
              @click="deleteComputeNode(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-drawer
      v-model="drawerFormVisible"
      :before-close="closeDrawer"
      :show-close="false"
      size="50%"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ drawerTitle }}</span>
          <div>
            <el-button @click="closeDrawer">取 消</el-button>
            <el-button type="primary" @click="enterDrawer">确 定</el-button>
          </div>
        </div>
      </template>
      <el-form ref="formRef" :model="form" :rules="rules" label-width="140px">
        <el-form-item label="节点名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入节点名称" />
        </el-form-item>
        <el-form-item label="区域">
          <el-input v-model="form.region" placeholder="例如: cn-north-1" />
        </el-form-item>
        <el-form-item label="CPU核心数">
          <el-input-number v-model="form.cpu" :min="0" />
        </el-form-item>
        <el-form-item label="内存(GB)">
          <el-input-number v-model="form.memory" :min="0" />
        </el-form-item>
        <el-form-item label="系统盘容量(GB)">
          <el-input-number v-model="form.systemDisk" :min="0" />
        </el-form-item>
        <el-form-item label="数据盘容量(GB)">
          <el-input-number v-model="form.dataDisk" :min="0" />
        </el-form-item>
        <el-form-item label="公网IP" prop="publicIP">
          <el-input v-model="form.publicIP" placeholder="请输入公网IP" />
        </el-form-item>
        <el-form-item label="内网IP" prop="privateIP">
          <el-input v-model="form.privateIP" placeholder="请输入内网IP" />
        </el-form-item>
        <el-form-item label="SSH端口">
          <el-input-number v-model="form.sshPort" :min="1" :max="65535" />
        </el-form-item>
        <el-form-item label="SSH用户名">
          <el-input v-model="form.username" placeholder="请输入SSH用户名" />
        </el-form-item>
        <el-form-item label="SSH密码">
          <el-input v-model="form.password" type="password" placeholder="请输入SSH密码" show-password />
        </el-form-item>
        <el-form-item label="显卡名称">
          <el-input v-model="form.gpuName" placeholder="例如: NVIDIA GeForce RTX 4090" />
        </el-form-item>
        <el-form-item label="显卡数量">
          <el-input-number v-model="form.gpuCount" :min="0" />
        </el-form-item>
        <el-form-item label="Docker连接地址">
          <el-input v-model="form.dockerEndpoint" placeholder="例如: tcp://localhost:2375" />
        </el-form-item>
        <el-form-item label="使用TLS">
          <el-switch v-model="form.useTLS" />
        </el-form-item>
        <el-form-item label="CA证书">
          <el-input v-model="form.tlsCaCert" type="textarea" :rows="3" placeholder="请输入CA证书内容" />
        </el-form-item>
        <el-form-item label="客户端证书">
          <el-input v-model="form.tlsCert" type="textarea" :rows="3" placeholder="请输入客户端证书内容" />
        </el-form-item>
        <el-form-item label="客户端私钥">
          <el-input v-model="form.tlsKey" type="textarea" :rows="3" placeholder="请输入客户端私钥内容" />
        </el-form-item>
        <el-form-item label="是否上架">
          <el-switch v-model="form.isPublished" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="form.remark" type="textarea" :rows="3" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createComputeNode,
  updateComputeNode,
  deleteComputeNode,
  getComputeNodeList
} from '@/api/k8sgpu.js'
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const drawerFormVisible = ref(false)
const drawerTitle = ref('新增算力节点')
const formRef = ref(null)

const form = reactive({
  ID: 0,
  name: '',
  region: '',
  cpu: 0,
  memory: 0,
  systemDisk: 0,
  dataDisk: 0,
  publicIP: '',
  privateIP: '',
  sshPort: 22,
  username: '',
  password: '',
  gpuName: '',
  gpuCount: 0,
  dockerEndpoint: '',
  useTLS: true,
  tlsCaCert: '',
  tlsCert: '',
  tlsKey: '',
  isPublished: true,
  remark: ''
})

const rules = {
  name: [{ required: true, message: '请输入节点名称', trigger: 'blur' }],
  publicIP: [{ required: true, message: '请输入公网IP', trigger: 'blur' }],
  privateIP: [{ required: true, message: '请输入内网IP', trigger: 'blur' }]
}

const getTableData = async () => {
  const res = await getComputeNodeList({
    page: page.value,
    pageSize: pageSize.value
  })
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
  }
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const openDrawer = () => {
  drawerTitle.value = '新增算力节点'
  Object.assign(form, {
    ID: 0,
    name: '',
    region: '',
    cpu: 0,
    memory: 0,
    systemDisk: 0,
    dataDisk: 0,
    publicIP: '',
    privateIP: '',
    sshPort: 22,
    username: '',
    password: '',
    gpuName: '',
    gpuCount: 0,
    dockerEndpoint: '',
    useTLS: true,
    tlsCaCert: '',
    tlsCert: '',
    tlsKey: '',
    isPublished: true,
    remark: ''
  })
  drawerFormVisible.value = true
}

const closeDrawer = () => {
  drawerFormVisible.value = false
  formRef.value?.resetFields()
}

const enterDrawer = async () => {
  formRef.value?.validate(async (valid) => {
    if (valid) {
      const data = { ...form }
      let res
      if (data.ID) {
        res = await updateComputeNode(data)
      } else {
        res = await createComputeNode(data)
      }
      if (res.code === 0) {
        ElMessage.success(data.ID ? '更新成功' : '创建成功')
        closeDrawer()
        getTableData()
      }
    }
  })
}

const updateComputeNode = (row) => {
  drawerTitle.value = '编辑算力节点'
  Object.assign(form, row)
  drawerFormVisible.value = true
}

const deleteComputeNode = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteComputeNode({ id: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

const togglePublished = async (row) => {
  const res = await updateComputeNode({ ...row, isPublished: row.isPublished })
  if (res.code === 0) {
    ElMessage.success('状态更新成功')
  } else {
    row.isPublished = !row.isPublished
  }
}

getTableData()
</script>

<style scoped lang="scss">
.flex {
  display: flex;
}
.justify-between {
  justify-content: space-between;
}
.items-center {
  align-items: center;
}
.text-lg {
  font-size: 1.125rem;
}
</style>
