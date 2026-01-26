<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDrawer">新增规格</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column align="left" label="ID" prop="ID" width="60" />
        <el-table-column align="left" label="规格名称" prop="name" width="150" />
        <el-table-column align="left" label="显卡型号" prop="gpuModel" width="150" />
        <el-table-column align="left" label="显卡数量" prop="gpuCount" width="100" />
        <el-table-column align="left" label="CPU核心数" prop="cpuCores" width="100" />
        <el-table-column align="left" label="内存(GB)" prop="memory" width="100" />
        <el-table-column align="left" label="系统盘(GB)" prop="systemDisk" width="100" />
        <el-table-column align="left" label="数据盘(GB)" prop="dataDisk" width="100" />
        <el-table-column align="left" label="价格/小时" prop="pricePerHour" width="120">
          <template #default="scope">
            ¥{{ scope.row.pricePerHour }}
          </template>
        </el-table-column>
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
              @click="updateProductSpec(scope.row)"
            >编辑</el-button>
            <el-button
              type="danger"
              link
              icon="delete"
              @click="deleteProductSpec(scope.row)"
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
      size="40%"
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
      <el-form ref="formRef" :model="form" :rules="rules" label-width="120px">
        <el-form-item label="规格名称" prop="name">
          <el-input v-model="form.name" placeholder="例如: RTX4090x2 高配版" />
        </el-form-item>
        <el-form-item label="显卡型号" prop="gpuModel">
          <el-input v-model="form.gpuModel" placeholder="例如: RTX4090" />
        </el-form-item>
        <el-form-item label="显卡数量">
          <el-input-number v-model="form.gpuCount" :min="0" />
        </el-form-item>
        <el-form-item label="CPU核心数">
          <el-input-number v-model="form.cpuCores" :min="0" />
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
        <el-form-item label="价格/小时">
          <el-input-number v-model="form.pricePerHour" :min="0" :precision="2" />
        </el-form-item>
        <el-form-item label="是否上架">
          <el-switch v-model="form.isPublished" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="form.remark" type="textarea" :rows="3" />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createProductSpec,
  updateProductSpec,
  deleteProductSpec,
  getProductSpecList
} from '@/api/k8sgpu.js'
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const drawerFormVisible = ref(false)
const drawerTitle = ref('新增产品规格')
const formRef = ref(null)

const form = reactive({
  ID: 0,
  name: '',
  gpuModel: '',
  gpuCount: 1,
  cpuCores: 0,
  memory: 0,
  systemDisk: 0,
  dataDisk: 0,
  pricePerHour: 0,
  isPublished: true,
  remark: ''
})

const rules = {
  name: [{ required: true, message: '请输入规格名称', trigger: 'blur' }],
  gpuModel: [{ required: true, message: '请输入显卡型号', trigger: 'blur' }]
}

const getTableData = async () => {
  const res = await getProductSpecList({
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
  drawerTitle.value = '新增产品规格'
  Object.assign(form, {
    ID: 0,
    name: '',
    gpuModel: '',
    gpuCount: 1,
    cpuCores: 0,
    memory: 0,
    systemDisk: 0,
    dataDisk: 0,
    pricePerHour: 0,
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
        res = await updateProductSpec(data)
      } else {
        res = await createProductSpec(data)
      }
      if (res.code === 0) {
        ElMessage.success(data.ID ? '更新成功' : '创建成功')
        closeDrawer()
        getTableData()
      }
    }
  })
}

const updateProductSpec = (row) => {
  drawerTitle.value = '编辑产品规格'
  Object.assign(form, row)
  drawerFormVisible.value = true
}

const deleteProductSpec = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteProductSpec({ id: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

const togglePublished = async (row) => {
  const res = await updateProductSpec({ ...row, isPublished: row.isPublished })
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
