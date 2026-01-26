<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDrawer">新增镜像</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="镜像名称" prop="name" width="150" />
        <el-table-column align="left" label="镜像地址" prop="address" min-width="200" show-overflow-tooltip />
        <el-table-column align="left" label="描述" prop="description" width="200" show-overflow-tooltip />
        <el-table-column align="left" label="来源" prop="source" width="120" />
        <el-table-column align="left" label="是否上架" width="100">
          <template #default="scope">
            <el-switch
              v-model="scope.row.isPublished"
              @change="togglePublished(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="备注" prop="remark" width="200" show-overflow-tooltip />
        <el-table-column align="left" label="操作" width="200">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              @click="updateImageRegistry(scope.row)"
            >编辑</el-button>
            <el-button
              type="danger"
              link
              icon="delete"
              @click="deleteImageRegistry(scope.row)"
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
      <el-form ref="formRef" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="镜像名称" prop="name">
          <el-input v-model="form.name" autocomplete="off" placeholder="请输入镜像名称" />
        </el-form-item>
        <el-form-item label="镜像地址" prop="address">
          <el-input v-model="form.address" autocomplete="off" placeholder="例如: nginx:latest" />
        </el-form-item>
        <el-form-item label="镜像描述">
          <el-input v-model="form.description" type="textarea" :rows="3" placeholder="请输入镜像描述" />
        </el-form-item>
        <el-form-item label="镜像来源">
          <el-input v-model="form.source" placeholder="例如: Docker Hub" />
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
  createImageRegistry,
  updateImageRegistry,
  deleteImageRegistry,
  getImageRegistryList
} from '@/api/k8sgpu.js'
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const drawerFormVisible = ref(false)
const drawerTitle = ref('新增镜像')
const formRef = ref(null)

const form = reactive({
  ID: 0,
  name: '',
  address: '',
  description: '',
  source: '',
  isPublished: true,
  remark: ''
})

const rules = {
  name: [{ required: true, message: '请输入镜像名称', trigger: 'blur' }],
  address: [{ required: true, message: '请输入镜像地址', trigger: 'blur' }]
}

// 获取列表
const getTableData = async () => {
  const res = await getImageRegistryList({
    page: page.value,
    pageSize: pageSize.value
  })
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
  }
}

// 分页
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 打开抽屉
const openDrawer = () => {
  drawerTitle.value = '新增镜像'
  Object.assign(form, {
    ID: 0,
    name: '',
    address: '',
    description: '',
    source: '',
    isPublished: true,
    remark: ''
  })
  drawerFormVisible.value = true
}

// 关闭抽屉
const closeDrawer = () => {
  drawerFormVisible.value = false
  formRef.value?.resetFields()
}

// 确定提交
const enterDrawer = async () => {
  formRef.value?.validate(async (valid) => {
    if (valid) {
      const data = { ...form }
      let res
      if (data.ID) {
        res = await updateImageRegistry(data)
      } else {
        res = await createImageRegistry(data)
      }
      if (res.code === 0) {
        ElMessage.success(data.ID ? '更新成功' : '创建成功')
        closeDrawer()
        getTableData()
      }
    }
  })
}

// 编辑
const updateImageRegistry = (row) => {
  drawerTitle.value = '编辑镜像'
  Object.assign(form, row)
  drawerFormVisible.value = true
}

// 删除
const deleteImageRegistry = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteImageRegistry({ id: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

// 切换上架状态
const togglePublished = async (row) => {
  const res = await updateImageRegistry({
    ID: row.ID,
    name: row.name,
    address: row.address,
    description: row.description,
    source: row.source,
    isPublished: row.isPublished,
    remark: row.remark
  })
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
