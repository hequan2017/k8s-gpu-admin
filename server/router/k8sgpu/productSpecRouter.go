package k8sgpu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProductSpecRouter struct{}

// InitProductSpecRouter 初始化产品规格路由
func (psr *ProductSpecRouter) InitProductSpecRouter(Router *gin.RouterGroup) {
	productSpecApi := v1.ApiGroupApp.K8sgpuApiGroup
	productSpecRouter := Router.Group("k8sgpu/productSpec").Use(middleware.OperationRecord())
	productSpecRouterWithoutRecord := Router.Group("k8sgpu/productSpec")
	{
		productSpecRouter.POST("", productSpecApi.ProductSpecApi.CreateProductSpec)   // 创建产品规格
		productSpecRouter.PUT("", productSpecApi.ProductSpecApi.UpdateProductSpec)    // 更新产品规格
		productSpecRouter.DELETE("", productSpecApi.ProductSpecApi.DeleteProductSpec) // 删除产品规格
	}
	{
		productSpecRouterWithoutRecord.GET("", productSpecApi.ProductSpecApi.GetProductSpec)         // 获取产品规格信息
		productSpecRouterWithoutRecord.GET("list", productSpecApi.ProductSpecApi.GetProductSpecList) // 获取产品规格列表
		productSpecRouterWithoutRecord.GET("published", productSpecApi.ProductSpecApi.GetPublishedProductSpecList) // 获取已上架的产品规格列表
	}
}
