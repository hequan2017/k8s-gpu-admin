package k8sgpu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ImageRegistryRouter struct{}

// InitImageRegistryRouter 初始化镜像库路由
func (irr *ImageRegistryRouter) InitImageRegistryRouter(Router *gin.RouterGroup) {
	imageRegistryApi := v1.ApiGroupApp.K8sgpuApiGroup
	imageRegistryRouter := Router.Group("k8sgpu/imageRegistry").Use(middleware.OperationRecord())
	imageRegistryRouterWithoutRecord := Router.Group("k8sgpu/imageRegistry")
	{
		imageRegistryRouter.POST("", imageRegistryApi.ImageRegistryApi.CreateImageRegistry)   // 创建镜像
		imageRegistryRouter.PUT("", imageRegistryApi.ImageRegistryApi.UpdateImageRegistry)    // 更新镜像
		imageRegistryRouter.DELETE("", imageRegistryApi.ImageRegistryApi.DeleteImageRegistry) // 删除镜像
	}
	{
		imageRegistryRouterWithoutRecord.GET("", imageRegistryApi.ImageRegistryApi.GetImageRegistry)         // 获取镜像信息
		imageRegistryRouterWithoutRecord.GET("list", imageRegistryApi.ImageRegistryApi.GetImageRegistryList) // 获取镜像列表
		imageRegistryRouterWithoutRecord.GET("published", imageRegistryApi.ImageRegistryApi.GetPublishedImageRegistryList) // 获取已上架的镜像列表
	}
}
