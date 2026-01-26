package k8sgpu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InstanceRouter struct{}

// InitInstanceRouter 初始化实例管理路由
func (ir *InstanceRouter) InitInstanceRouter(Router *gin.RouterGroup) {
	instanceApi := v1.ApiGroupApp.K8sgpuApiGroup
	instanceRouter := Router.Group("k8sgpu/instance").Use(middleware.OperationRecord())
	instanceRouterWithoutRecord := Router.Group("k8sgpu/instance")
	{
		instanceRouter.POST("", instanceApi.InstanceApi.CreateInstance)   // 创建实例
		instanceRouter.DELETE("", instanceApi.InstanceApi.DeleteInstance) // 删除实例
	}
	{
		instanceRouterWithoutRecord.GET("", instanceApi.InstanceApi.GetInstance)         // 获取实例信息
		instanceRouterWithoutRecord.GET("list", instanceApi.InstanceApi.GetInstanceList) // 获取实例列表
	}
	{
		instanceRouter.POST("containerAction", instanceApi.InstanceApi.ContainerAction)         // 容器操作
		instanceRouter.POST("containerLogs", instanceApi.InstanceApi.GetContainerLogs)          // 获取容器日志
		instanceRouter.POST("updateStatus", instanceApi.InstanceApi.UpdateInstanceStatus)       // 更新实例状态
	}
}
