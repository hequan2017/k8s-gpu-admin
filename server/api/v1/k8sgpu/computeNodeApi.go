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

type ComputeNodeApi struct{}

// CreateComputeNode 创建算力节点
// @Tags      ComputeNode
// @Summary   创建算力节点
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      k8sgpuReq.ComputeNodeAdd  true  "节点名称, 公网IP, 内网IP"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /k8sgpu/computeNode [post]
func (cna *ComputeNodeApi) CreateComputeNode(c *gin.Context) {
	var req k8sgpuReq.ComputeNodeAdd
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.K8sgpuServiceGroup.ComputeNodeServiceApp.CreateComputeNode(req)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteComputeNode 删除算力节点
// @Tags      ComputeNode
// @Summary   删除算力节点
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                    true  "节点ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /k8sgpu/computeNode [delete]
func (cna *ComputeNodeApi) DeleteComputeNode(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.K8sgpuServiceGroup.ComputeNodeServiceApp.DeleteComputeNode(req.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateComputeNode 更新算力节点
// @Tags      ComputeNode
// @Summary   更新算力节点
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      k8sgpuReq.ComputeNodeUpdate  true  "节点ID, 节点信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /k8sgpu/computeNode [put]
func (cna *ComputeNodeApi) UpdateComputeNode(c *gin.Context) {
	var req k8sgpuReq.ComputeNodeUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.K8sgpuServiceGroup.ComputeNodeServiceApp.UpdateComputeNode(req)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetComputeNode 获取算力节点信息
// @Tags      ComputeNode
// @Summary   获取算力节点信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                                   true  "节点ID"
// @Success   200   {object}  response.Response{data=k8sgpu.ComputeNode,msg=string}  "获取成功"
// @Router    /k8sgpu/computeNode [get]
func (cna *ComputeNodeApi) GetComputeNode(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := service.ServiceGroupApp.K8sgpuServiceGroup.ComputeNodeServiceApp.GetComputeNode(req.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}
	response.OkWithData(data, c)
}

// GetComputeNodeList 分页获取算力节点列表
// @Tags      ComputeNode
// @Summary   分页获取算力节点列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     k8sgpuReq.ComputeNodeSearch                   true  "分页信息"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /k8sgpu/computeNodeList [get]
func (cna *ComputeNodeApi) GetComputeNodeList(c *gin.Context) {
	var info k8sgpuReq.ComputeNodeSearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := service.ServiceGroupApp.K8sgpuServiceGroup.ComputeNodeServiceApp.GetComputeNodeList(info)
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

// MatchAvailableNodes 匹配可用节点
// @Tags      ComputeNode
// @Summary   匹配可用节点(根据产品规格)
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      k8sgpuReq.MatchNodeRequest  true  "产品规格ID"
// @Success   200   {object}  response.Response{data=[]k8sgpu.ComputeNode,msg=string}  "获取成功"
// @Router    /k8sgpu/matchAvailableNodes [post]
func (cna *ComputeNodeApi) MatchAvailableNodes(c *gin.Context) {
	var req k8sgpuReq.MatchNodeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	nodes, err := service.ServiceGroupApp.K8sgpuServiceGroup.ComputeNodeServiceApp.MatchAvailableNodes(req.SpecID)
	if err != nil {
		global.GVA_LOG.Error("匹配节点失败!", zap.Error(err))
		response.FailWithMessage("匹配节点失败: "+err.Error(), c)
		return
	}
	response.OkWithData(nodes, c)
}
