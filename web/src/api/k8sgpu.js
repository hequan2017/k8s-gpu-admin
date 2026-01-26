import service from '@/utils/request'

// ==================== 镜像库相关API ====================

// @Tags ImageRegistry
// @Summary 创建镜像
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body k8sgpu.ImageRegistryAdd true "镜像名称, 镜像地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /k8sgpu/imageRegistry [post]
export const createImageRegistry = (data) => {
  return service({
    url: '/k8sgpu/imageRegistry',
    method: 'post',
    data
  })
}

// @Tags ImageRegistry
// @Summary 更新镜像
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body k8sgpu.ImageRegistryUpdate true "镜像ID, 镜像信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sgpu/imageRegistry [put]
export const updateImageRegistry = (data) => {
  return service({
    url: '/k8sgpu/imageRegistry',
    method: 'put',
    data
  })
}

// @Tags ImageRegistry
// @Summary 删除镜像
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "镜像ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sgpu/imageRegistry [delete]
export const deleteImageRegistry = (data) => {
  return service({
    url: '/k8sgpu/imageRegistry',
    method: 'delete',
    data
  })
}

// @Tags ImageRegistry
// @Summary 获取镜像信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetById true "镜像ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sgpu/imageRegistry [get]
export const getImageRegistry = (params) => {
  return service({
    url: '/k8sgpu/imageRegistry',
    method: 'get',
    params
  })
}

// @Tags ImageRegistry
// @Summary 分页获取镜像列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query k8sgpu.ImageRegistrySearch true "分页信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sgpu/imageRegistry/list [get]
export const getImageRegistryList = (params) => {
  return service({
    url: '/k8sgpu/imageRegistry/list',
    method: 'get',
    params
  })
}

// @Tags ImageRegistry
// @Summary 获取已上架的镜像列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sgpu/imageRegistry/published [get]
export const getPublishedImageRegistryList = () => {
  return service({
    url: '/k8sgpu/imageRegistry/published',
    method: 'get'
  })
}

// ==================== 算力节点相关API ====================

// @Tags ComputeNode
// @Summary 创建算力节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body k8sgpu.ComputeNodeAdd true "节点名称, 公网IP, 内网IP"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /k8sgpu/computeNode [post]
export const createComputeNode = (data) => {
  return service({
    url: '/k8sgpu/computeNode',
    method: 'post',
    data
  })
}

// @Tags ComputeNode
// @Summary 更新算力节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body k8sgpu.ComputeNodeUpdate true "节点ID, 节点信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sgpu/computeNode [put]
export const updateComputeNode = (data) => {
  return service({
    url: '/k8sgpu/computeNode',
    method: 'put',
    data
  })
}

// @Tags ComputeNode
// @Summary 删除算力节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "节点ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sgpu/computeNode [delete]
export const deleteComputeNode = (data) => {
  return service({
    url: '/k8sgpu/computeNode',
    method: 'delete',
    data
  })
}

// @Tags ComputeNode
// @Summary 获取算力节点信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetById true "节点ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sgpu/computeNode [get]
export const getComputeNode = (params) => {
  return service({
    url: '/k8sgpu/computeNode',
    method: 'get',
    params
  })
}

// @Tags ComputeNode
// @Summary 分页获取算力节点列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query k8sgpu.ComputeNodeSearch true "分页信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sgpu/computeNode/list [get]
export const getComputeNodeList = (params) => {
  return service({
    url: '/k8sgpu/computeNode/list',
    method: 'get',
    params
  })
}

// @Tags ComputeNode
// @Summary 匹配可用节点(根据产品规格)
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body k8sgpu.MatchNodeRequest true "产品规格ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sgpu/computeNode/matchAvailableNodes [post]
export const matchAvailableNodes = (data) => {
  return service({
    url: '/k8sgpu/computeNode/matchAvailableNodes',
    method: 'post',
    data
  })
}

// ==================== 产品规格相关API ====================

// @Tags ProductSpec
// @Summary 创建产品规格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body k8sgpu.ProductSpecAdd true "规格名称, 显卡型号"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /k8sgpu/productSpec [post]
export const createProductSpec = (data) => {
  return service({
    url: '/k8sgpu/productSpec',
    method: 'post',
    data
  })
}

// @Tags ProductSpec
// @Summary 更新产品规格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body k8sgpu.ProductSpecUpdate true "规格ID, 规格信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sgpu/productSpec [put]
export const updateProductSpec = (data) => {
  return service({
    url: '/k8sgpu/productSpec',
    method: 'put',
    data
  })
}

// @Tags ProductSpec
// @Summary 删除产品规格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "规格ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sgpu/productSpec [delete]
export const deleteProductSpec = (data) => {
  return service({
    url: '/k8sgpu/productSpec',
    method: 'delete',
    data
  })
}

// @Tags ProductSpec
// @Summary 获取产品规格信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetById true "规格ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sgpu/productSpec [get]
export const getProductSpec = (params) => {
  return service({
    url: '/k8sgpu/productSpec',
    method: 'get',
    params
  })
}

// @Tags ProductSpec
// @Summary 分页获取产品规格列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query k8sgpu.ProductSpecSearch true "分页信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sgpu/productSpec/list [get]
export const getProductSpecList = (params) => {
  return service({
    url: '/k8sgpu/productSpec/list',
    method: 'get',
    params
  })
}

// @Tags ProductSpec
// @Summary 获取已上架的产品规格列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sgpu/productSpec/published [get]
export const getPublishedProductSpecList = () => {
  return service({
    url: '/k8sgpu/productSpec/published',
    method: 'get'
  })
}

// ==================== 实例管理相关API ====================

// @Tags Instance
// @Summary 创建实例
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body k8sgpu.InstanceAdd true "实例名称, 镜像ID, 规格ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /k8sgpu/instance [post]
export const createInstance = (data) => {
  return service({
    url: '/k8sgpu/instance',
    method: 'post',
    data
  })
}

// @Tags Instance
// @Summary 删除实例
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "实例ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sgpu/instance [delete]
export const deleteInstance = (data) => {
  return service({
    url: '/k8sgpu/instance',
    method: 'delete',
    data
  })
}

// @Tags Instance
// @Summary 获取实例信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetById true "实例ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sgpu/instance [get]
export const getInstance = (params) => {
  return service({
    url: '/k8sgpu/instance',
    method: 'get',
    params
  })
}

// @Tags Instance
// @Summary 分页获取实例列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query k8sgpu.InstanceSearch true "分页信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sgpu/instance/list [get]
export const getInstanceList = (params) => {
  return service({
    url: '/k8sgpu/instance/list',
    method: 'get',
    params
  })
}

// @Tags Instance
// @Summary 容器操作(启动/停止/重启)
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body k8sgpu.ContainerActionRequest true "实例ID, 操作类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /k8sgpu/instance/containerAction [post]
export const containerAction = (data) => {
  return service({
    url: '/k8sgpu/instance/containerAction',
    method: 'post',
    data
  })
}

// @Tags Instance
// @Summary 获取容器日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "实例ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sgpu/instance/containerLogs [post]
export const getContainerLogs = (data) => {
  return service({
    url: '/k8sgpu/instance/containerLogs',
    method: 'post',
    data
  })
}

// @Tags Instance
// @Summary 更新实例状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "实例ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sgpu/instance/updateStatus [post]
export const updateInstanceStatus = (data) => {
  return service({
    url: '/k8sgpu/instance/updateStatus',
    method: 'post',
    data
  })
}
