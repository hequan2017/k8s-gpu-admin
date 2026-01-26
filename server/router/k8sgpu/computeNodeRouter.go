package k8sgpu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ComputeNodeRouter struct{}

// InitComputeNodeRouter 初始化算力节点路由
func (cnr *ComputeNodeRouter) InitComputeNodeRouter(Router *gin.RouterGroup) {
	computeNodeApi := v1.ApiGroupApp.K8sgpuApiGroup
	computeNodeRouter := Router.Group("k8sgpu/computeNode").Use(middleware.OperationRecord())
	computeNodeRouterWithoutRecord := Router.Group("k8sgpu/computeNode")
	{
		computeNodeRouter.POST("", computeNodeApi.ComputeNodeApi.CreateComputeNode)   // 创建算力节点
		computeNodeRouter.PUT("", computeNodeApi.ComputeNodeApi.UpdateComputeNode)    // 更新算力节点
		computeNodeRouter.DELETE("", computeNodeApi.ComputeNodeApi.DeleteComputeNode) // 删除算力节点
	}
	{
		computeNodeRouterWithoutRecord.GET("", computeNodeApi.ComputeNodeApi.GetComputeNode)         // 获取算力节点信息
		computeNodeRouterWithoutRecord.GET("list", computeNodeApi.ComputeNodeApi.GetComputeNodeList) // 获取算力节点列表
	}
	{
		computeNodeRouterWithoutRecord.POST("matchAvailableNodes", computeNodeApi.ComputeNodeApi.MatchAvailableNodes) // 匹配可用节点
	}
}
