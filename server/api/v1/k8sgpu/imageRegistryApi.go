package k8sgpu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/k8sgpu"
	k8sgpuReq "github.com/flipped-aurora/gin-vue-admin/server/model/k8sgpu/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ImageRegistryApi struct{}

// CreateImageRegistry 创建镜像
// @Tags      ImageRegistry
// @Summary   创建镜像
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      k8sgpuReq.ImageRegistryAdd  true  "镜像名称, 镜像地址"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /k8sgpu/imageRegistry [post]
func (ira *ImageRegistryApi) CreateImageRegistry(c *gin.Context) {
	var req k8sgpuReq.ImageRegistryAdd
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.K8sgpuServiceGroup.ImageRegistryServiceApp.CreateImageRegistry(req)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteImageRegistry 删除镜像
// @Tags      ImageRegistry
// @Summary   删除镜像
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                    true  "镜像ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /k8sgpu/imageRegistry [delete]
func (ira *ImageRegistryApi) DeleteImageRegistry(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.K8sgpuServiceGroup.ImageRegistryServiceApp.DeleteImageRegistry(req.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateImageRegistry 更新镜像
// @Tags      ImageRegistry
// @Summary   更新镜像
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      k8sgpuReq.ImageRegistryUpdate  true  "镜像ID, 镜像信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /k8sgpu/imageRegistry [put]
func (ira *ImageRegistryApi) UpdateImageRegistry(c *gin.Context) {
	var req k8sgpuReq.ImageRegistryUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.K8sgpuServiceGroup.ImageRegistryServiceApp.UpdateImageRegistry(req)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetImageRegistry 获取镜像信息
// @Tags      ImageRegistry
// @Summary   获取镜像信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                                   true  "镜像ID"
// @Success   200   {object}  response.Response{data=k8sgpu.ImageRegistry,msg=string}  "获取成功"
// @Router    /k8sgpu/imageRegistry [get]
func (ira *ImageRegistryApi) GetImageRegistry(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := service.ServiceGroupApp.K8sgpuServiceGroup.ImageRegistryServiceApp.GetImageRegistry(req.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}
	response.OkWithData(data, c)
}

// GetImageRegistryList 分页获取镜像列表
// @Tags      ImageRegistry
// @Summary   分页获取镜像列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     k8sgpuReq.ImageRegistrySearch                   true  "分页信息"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /k8sgpu/imageRegistryList [get]
func (ira *ImageRegistryApi) GetImageRegistryList(c *gin.Context) {
	var info k8sgpuReq.ImageRegistrySearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := service.ServiceGroupApp.K8sgpuServiceGroup.ImageRegistryServiceApp.GetImageRegistryList(info)
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

// GetPublishedImageRegistryList 获取已上架的镜像列表(用于实例创建选择)
// @Tags      ImageRegistry
// @Summary   获取已上架的镜像列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=[]k8sgpu.ImageRegistry,msg=string}  "获取成功"
// @Router    /k8sgpu/publishedImageRegistryList [get]
func (ira *ImageRegistryApi) GetPublishedImageRegistryList(c *gin.Context) {
	list, err := service.ServiceGroupApp.K8sgpuServiceGroup.ImageRegistryServiceApp.GetPublishedImageRegistryList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}
