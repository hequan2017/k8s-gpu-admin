<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openCreateDialog">新增实例</el-button>
        <el-button icon="refresh" @click="refreshList">刷新状态</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column align="left" label="ID" prop="ID" width="60" />
        <el-table-column align="left" label="实例名称" prop="instanceName" width="150" />
        <el-table-column align="left" label="镜像" prop="image.name" width="150" />
        <el-table-column align="left" label="规格" width="200">
          <template #default="scope">
            {{ scope.row.spec.name }}
            <span class="text-gray-500 text-sm">
              {{ scope.row.spec.gpuModel }} x{{ scope.row.spec.gpuCount }}
            </span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="节点" prop="node.name" width="120" />
        <el-table-column align="left" label="容器ID" prop="containerId" width="120" show-overflow-tooltip />
        <el-table-column align="left" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.containerStatus)">
              {{ scope.row.containerStatus || 'unknown' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" width="300" fixed="right">
          <template #default="scope">
            <el-button
              type="success"
              link
              icon="video-play"
              :disabled="!scope.row.containerId || scope.row.containerStatus === 'running'"
              @click="containerAction(scope.row, 'start')"
            >启动</el-button>
            <el-button
              type="warning"
              link
              icon="video-pause"
              :disabled="!scope.row.containerId || scope.row.containerStatus === 'stopped'"
              @click="containerAction(scope.row, 'stop')"
            >停止</el-button>
            <el-button
              type="primary"
              link
              icon="refresh"
              :disabled="!scope.row.containerId"
              @click="containerAction(scope.row, 'restart')"
            >重启</el-button>
            <el-button
              type="info"
              link
              icon="document"
              :disabled="!scope.row.containerId"
              @click="showLogs(scope.row)"
            >日志</el-button>
            <el-button
              type="danger"
              link
              icon="delete"
              @click="deleteInstance(scope.row)"
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

    <!-- 新增实例对话框 -->
    <el-dialog
      v-model="createDialogVisible"
      title="新增实例"
      width="600px"
      :close-on-click-modal="false"
    >
      <el-steps :active="currentStep" align-center finish-status="success">
        <el-step title="选择镜像" />
        <el-step title="选择规格" />
        <el-step title="确认信息" />
      </el-steps>

      <div class="step-content" v-if="currentStep === 0">
        <el-form :model="createForm" label-width="100px">
          <el-form-item label="选择镜像">
            <el-select v-model="createForm.imageId" placeholder="请选择镜像" style="width: 100%">
              <el-option
                v-for="item in publishedImages"
                :key="item.ID"
                :label="item.name"
                :value="item.ID"
              >
                <span>{{ item.name }}</span>
                <span style="float: right; color: var(--el-text-color-secondary); font-size: 13px">
                  {{ item.address }}
                </span>
              </el-option>
            </el-select>
          </el-form-item>
        </el-form>
      </div>

      <div class="step-content" v-if="currentStep === 1">
        <el-form :model="createForm" label-width="100px">
          <el-form-item label="选择规格">
            <el-select v-model="createForm.specId" placeholder="请选择规格" style="width: 100%" @change="onSpecChange">
              <el-option
                v-for="item in publishedSpecs"
                :key="item.ID"
                :label="getSpecLabel(item)"
                :value="item.ID"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="可用节点" v-if="matchedNodes.length > 0">
            <el-select v-model="createForm.nodeId" placeholder="系统将自动选择合适节点" style="width: 100%">
              <el-option
                v-for="item in matchedNodes"
                :key="item.ID"
                :label="item.name"
                :value="item.ID"
              >
                <span>{{ item.name }}</span>
                <span style="float: right; color: var(--el-text-color-secondary); font-size: 13px">
                  GPU: {{ item.gpuName }} x{{ item.gpuCount }}
                </span>
              </el-option>
            </el-select>
          </el-form-item>
          <el-alert
            v-if="matchedNodes.length === 0 && createForm.specId"
            title="没有可用的算力节点"
            type="warning"
            :closable="false"
          />
        </el-form>
      </div>

      <div class="step-content" v-if="currentStep === 2">
        <el-descriptions :column="1" border>
          <el-descriptions-item label="实例名称">{{ createForm.instanceName }}</el-descriptions-item>
          <el-descriptions-item label="镜像">{{ selectedImage?.name }}</el-descriptions-item>
          <el-descriptions-item label="规格">{{ getSpecLabel(selectedSpec) }}</el-descriptions-item>
          <el-descriptions-item label="节点">{{ matchedNodes[0]?.name || '自动匹配' }}</el-descriptions-item>
        </el-descriptions>
        <el-form :model="createForm" label-width="100px" style="margin-top: 20px">
          <el-form-item label="实例名称">
            <el-input v-model="createForm.instanceName" placeholder="请输入实例名称" />
          </el-form-item>
          <el-form-item label="备注">
            <el-input v-model="createForm.remark" type="textarea" :rows="3" />
          </el-form-item>
        </el-form>
      </div>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="createDialogVisible = false">取消</el-button>
          <el-button v-if="currentStep > 0" @click="currentStep--">上一步</el-button>
          <el-button v-if="currentStep < 2" type="primary" @click="nextStep">下一步</el-button>
          <el-button v-if="currentStep === 2" type="primary" @click="submitCreate">创建</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 日志对话框 -->
    <el-dialog
      v-model="logsDialogVisible"
      title="容器日志"
      width="800px"
    >
      <div class="logs-content">
        <pre>{{ logsContent }}</pre>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createInstance,
  deleteInstance,
  getInstanceList,
  containerAction,
  getContainerLogs,
  getPublishedImageRegistryList,
  getPublishedProductSpecList,
  matchAvailableNodes,
  updateInstanceStatus
} from '@/api/k8sgpu.js'
import { ref, reactive, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)

const createDialogVisible = ref(false)
const currentStep = ref(0)
const publishedImages = ref([])
const publishedSpecs = ref([])
const matchedNodes = ref([])

const logsDialogVisible = ref(false)
const logsContent = ref('')

const createForm = reactive({
  instanceName: '',
  imageId: null,
  specId: null,
  nodeId: null,
  remark: ''
})

const selectedImage = computed(() => publishedImages.value.find(i => i.ID === createForm.imageId))
const selectedSpec = computed(() => publishedSpecs.value.find(s => s.ID === createForm.specId))

const getTableData = async () => {
  const res = await getInstanceList({
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

const refreshList = async () => {
  // 更新所有实例状态
  for (const item of tableData.value) {
    if (item.ID) {
      await updateInstanceStatus({ id: item.ID })
    }
  }
  getTableData()
  ElMessage.success('状态已刷新')
}

const getSpecLabel = (spec) => {
  if (!spec) return ''
  return `${spec.name} - ${spec.gpuModel} x${spec.gpuCount} / ${spec.cpuCores}核 / ${spec.memory}GB`
}

const getStatusType = (status) => {
  const map = {
    running: 'success',
    stopped: 'info',
    creating: 'warning',
    removed: 'danger'
  }
  return map[status] || 'info'
}

const openCreateDialog = async () => {
  currentStep.value = 0
  Object.assign(createForm, {
    instanceName: '',
    imageId: null,
    specId: null,
    nodeId: null,
    remark: ''
  })
  matchedNodes.value = []

  // 加载镜像和规格列表
  const [imagesRes, specsRes] = await Promise.all([
    getPublishedImageRegistryList(),
    getPublishedProductSpecList()
  ])
  if (imagesRes.code === 0) publishedImages.value = imagesRes.data
  if (specsRes.code === 0) publishedSpecs.value = specsRes.data

  createDialogVisible.value = true
}

const nextStep = async () => {
  if (currentStep.value === 0) {
    if (!createForm.imageId) {
      ElMessage.warning('请选择镜像')
      return
    }
    currentStep.value++
  } else if (currentStep.value === 1) {
    if (!createForm.specId) {
      ElMessage.warning('请选择规格')
      return
    }
    // 匹配可用节点
    const res = await matchAvailableNodes({ specId: createForm.specId })
    if (res.code === 0) {
      matchedNodes.value = res.data
      if (matchedNodes.value.length > 0) {
        createForm.nodeId = matchedNodes.value[0].ID
      }
    }
    currentStep.value++
  }
}

const onSpecChange = () => {
  createForm.nodeId = null
  matchedNodes.value = []
}

const submitCreate = async () => {
  if (!createForm.instanceName) {
    ElMessage.warning('请输入实例名称')
    return
  }

  const res = await createInstance(createForm)
  if (res.code === 0) {
    ElMessage.success('创建成功')
    createDialogVisible.value = false
    getTableData()
  }
}

const containerAction = async (row, action) => {
  const actionName = { start: '启动', stop: '停止', restart: '重启' }[action]
  const res = await containerAction({ id: row.ID, action })
  if (res.code === 0) {
    ElMessage.success(`${actionName}成功`)
    getTableData()
  }
}

const showLogs = async (row) => {
  logsDialogVisible.value = true
  logsContent.value = '加载中...'
  const res = await getContainerLogs({ id: row.ID })
  if (res.code === 0) {
    logsContent.value = res.data || '暂无日志'
  } else {
    logsContent.value = '获取日志失败'
  }
}

const deleteInstance = (row) => {
  ElMessageBox.confirm('确定要删除此实例吗？删除后将同时删除容器。', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteInstance({ id: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

getTableData()
</script>

<style scoped lang="scss">
.step-content {
  margin-top: 30px;
  min-height: 200px;
}

.text-gray-500 {
  color: #6b7280;
}

.text-sm {
  font-size: 0.875rem;
}

.logs-content {
  background: #1e1e1e;
  color: #d4d4d4;
  padding: 15px;
  border-radius: 4px;
  max-height: 400px;
  overflow-y: auto;

  pre {
    margin: 0;
    font-family: 'Courier New', monospace;
    font-size: 12px;
    white-space: pre-wrap;
    word-wrap: break-word;
  }
}
</style>
