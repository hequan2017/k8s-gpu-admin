package k8sgpu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/k8sgpu"
	k8sgpuReq "github.com/flipped-aurora/gin-vue-admin/server/model/k8sgpu/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type InstanceApi struct{}

// CreateInstance 创建实例
// @Tags      Instance
// @Summary   创建实例
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      k8sgpuReq.InstanceAdd  true  "实例名称, 镜像ID, 规格ID"
// @Success   200   {object}  response.Response{data=k8sgpu.Instance,msg=string}  "创建成功"
// @Router    /k8sgpu/instance [post]
func (ia *InstanceApi) CreateInstance(c *gin.Context) {
	var req k8sgpuReq.InstanceAdd
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	isAdmin := utils.GetUserAuthorityId(c) == 888 // 假设888是管理员权限ID

	instance, err := service.ServiceGroupApp.K8sgpuServiceGroup.InstanceServiceApp.CreateInstance(req, userID)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(instance, "创建成功", c)
}

// DeleteInstance 删除实例
// @Tags      Instance
// @Summary   删除实例
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                    true  "实例ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /k8sgpu/instance [delete]
func (ia *InstanceApi) DeleteInstance(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	isAdmin := utils.GetUserAuthorityId(c) == 888 // 假设888是管理员权限ID

	err = service.ServiceGroupApp.K8sgpuServiceGroup.InstanceServiceApp.DeleteInstance(req.ID, userID, isAdmin)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetInstance 获取实例信息
// @Tags      Instance
// @Summary   获取实例信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                                   true  "实例ID"
// @Success   200   {object}  response.Response{data=k8sgpu.Instance,msg=string}  "获取成功"
// @Router    /k8sgpu/instance [get]
func (ia *InstanceApi) GetInstance(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	isAdmin := utils.GetUserAuthorityId(c) == 888

	data, err := service.ServiceGroupApp.K8sgpuServiceGroup.InstanceServiceApp.GetInstance(req.ID, userID, isAdmin)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}
	response.OkWithData(data, c)
}

// GetInstanceList 分页获取实例列表
// @Tags      Instance
// @Summary   分页获取实例列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     k8sgpuReq.InstanceSearch                   true  "分页信息"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /k8sgpu/instanceList [get]
func (ia *InstanceApi) GetInstanceList(c *gin.Context) {
	var info k8sgpuReq.InstanceSearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	isAdmin := utils.GetUserAuthorityId(c) == 888

	list, total, err := service.ServiceGroupApp.K8sgpuServiceGroup.InstanceServiceApp.GetInstanceList(info, userID, isAdmin)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "获取成功", c)
}

// ContainerAction 容器操作
// @Tags      Instance
// @Summary   容器操作(启动/停止/重启)
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      k8sgpuReq.ContainerActionRequest  true  "实例ID, 操作类型"
// @Success   200   {object}  response.Response{msg=string}  "操作成功"
// @Router    /k8sgpu/containerAction [post]
func (ia *InstanceApi) ContainerAction(c *gin.Context) {
	var req k8sgpuReq.ContainerActionRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	isAdmin := utils.GetUserAuthorityId(c) == 888

	err = service.ServiceGroupApp.K8sgpuServiceGroup.InstanceServiceApp.ContainerAction(req.ID, req.Action, userID, isAdmin)
	if err != nil {
		global.GVA_LOG.Error("操作失败!", zap.Error(err))
		response.FailWithMessage("操作失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("操作成功", c)
}

// GetContainerLogs 获取容器日志
// @Tags      Instance
// @Summary   获取容器日志
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById  true  "实例ID"
// @Success   200   {object}  response.Response{data=string,msg=string}  "获取成功"
// @Router    /k8sgpu/containerLogs [post]
func (ia *InstanceApi) GetContainerLogs(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	isAdmin := utils.GetUserAuthorityId(c) == 888

	logs, err := service.ServiceGroupApp.K8sgpuServiceGroup.InstanceServiceApp.GetContainerLogs(req.ID, userID, isAdmin)
	if err != nil {
		global.GVA_LOG.Error("获取日志失败!", zap.Error(err))
		response.FailWithMessage("获取日志失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(logs, "获取成功", c)
}

// UpdateInstanceStatus 更新实例状态
// @Tags      Instance
// @Summary   更新实例状态
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById  true  "实例ID"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /k8sgpu/updateInstanceStatus [post]
func (ia *InstanceApi) UpdateInstanceStatus(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.K8sgpuServiceGroup.InstanceServiceApp.UpdateInstanceStatus(req.ID)
	if err != nil {
		global.GVA_LOG.Error("更新状态失败!", zap.Error(err))
		response.FailWithMessage("更新状态失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
