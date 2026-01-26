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

type ProductSpecApi struct{}

// CreateProductSpec 创建产品规格
// @Tags      ProductSpec
// @Summary   创建产品规格
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      k8sgpuReq.ProductSpecAdd  true  "规格名称, 显卡型号"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /k8sgpu/productSpec [post]
func (psa *ProductSpecApi) CreateProductSpec(c *gin.Context) {
	var req k8sgpuReq.ProductSpecAdd
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.K8sgpuServiceGroup.ProductSpecServiceApp.CreateProductSpec(req)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteProductSpec 删除产品规格
// @Tags      ProductSpec
// @Summary   删除产品规格
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                    true  "规格ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /k8sgpu/productSpec [delete]
func (psa *ProductSpecApi) DeleteProductSpec(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.K8sgpuServiceGroup.ProductSpecServiceApp.DeleteProductSpec(req.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateProductSpec 更新产品规格
// @Tags      ProductSpec
// @Summary   更新产品规格
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      k8sgpuReq.ProductSpecUpdate  true  "规格ID, 规格信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /k8sgpu/productSpec [put]
func (psa *ProductSpecApi) UpdateProductSpec(c *gin.Context) {
	var req k8sgpuReq.ProductSpecUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.K8sgpuServiceGroup.ProductSpecServiceApp.UpdateProductSpec(req)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetProductSpec 获取产品规格信息
// @Tags      ProductSpec
// @Summary   获取产品规格信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                                   true  "规格ID"
// @Success   200   {object}  response.Response{data=k8sgpu.ProductSpec,msg=string}  "获取成功"
// @Router    /k8sgpu/productSpec [get]
func (psa *ProductSpecApi) GetProductSpec(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := service.ServiceGroupApp.K8sgpuServiceGroup.ProductSpecServiceApp.GetProductSpec(req.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}
	response.OkWithData(data, c)
}

// GetProductSpecList 分页获取产品规格列表
// @Tags      ProductSpec
// @Summary   分页获取产品规格列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     k8sgpuReq.ProductSpecSearch                   true  "分页信息"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /k8sgpu/productSpecList [get]
func (psa *ProductSpecApi) GetProductSpecList(c *gin.Context) {
	var info k8sgpuReq.ProductSpecSearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := service.ServiceGroupApp.K8sgpuServiceGroup.ProductSpecServiceApp.GetProductSpecList(info)
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

// GetPublishedProductSpecList 获取已上架的产品规格列表(用于实例创建选择)
// @Tags      ProductSpec
// @Summary   获取已上架的产品规格列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=[]k8sgpu.ProductSpec,msg=string}  "获取成功"
// @Router    /k8sgpu/publishedProductSpecList [get]
func (psa *ProductSpecApi) GetPublishedProductSpecList(c *gin.Context) {
	list, err := service.ServiceGroupApp.K8sgpuServiceGroup.ProductSpecServiceApp.GetPublishedProductSpecList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}
