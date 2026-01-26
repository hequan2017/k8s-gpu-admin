package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// ImageRegistrySearch 镜像库搜索结构体
type ImageRegistrySearch struct {
	Name        string `json:"name" form:"name" gorm:"like;comment:镜像名称"`       // 镜像名称(模糊查询)
	Address     string `json:"address" form:"address" gorm:"like;comment:镜像地址"`  // 镜像地址(模糊查询)
	Source      string `json:"source" form:"source" gorm:"comment:镜像来源"`         // 镜像来源
	IsPublished *bool  `json:"isPublished" form:"isPublished" gorm:"comment:是否上架"` // 是否上架
	request.PageInfo
}

// ImageRegistryAdd 镜像库添加请求
type ImageRegistryAdd struct {
	Name        string `json:"name" binding:"required" form:"name" gorm:"comment:镜像名称"`               // 镜像名称(必填)
	Address     string `json:"address" binding:"required" form:"address" gorm:"comment:镜像地址"`          // 镜像地址(必填)
	Description string `json:"description" form:"description" gorm:"comment:镜像描述"`                     // 镜像描述
	Source      string `json:"source" form:"source" gorm:"comment:镜像来源"`                              // 镜像来源
	IsPublished bool   `json:"isPublished" form:"isPublished" gorm:"default:true;comment:是否上架"`      // 是否上架(默认上架)
	Remark      string `json:"remark" form:"remark" gorm:"comment:备注"`                                 // 备注
}

// ImageRegistryUpdate 镜像库更新请求
type ImageRegistryUpdate struct {
	ID          uint    `json:"id" binding:"required" form:"id" gorm:"primarykey;comment:主键ID"` // 主键ID
	Name        string  `json:"name" binding:"required" form:"name" gorm:"comment:镜像名称"`        // 镜像名称(必填)
	Address     string  `json:"address" binding:"required" form:"address" gorm:"comment:镜像地址"`   // 镜像地址(必填)
	Description string  `json:"description" form:"description" gorm:"comment:镜像描述"`              // 镜像描述
	Source      string  `json:"source" form:"source" gorm:"comment:镜像来源"`                       // 镜像来源
	IsPublished *bool   `json:"isPublished" form:"isPublished" gorm:"comment:是否上架"`             // 是否上架(使用指针以区分零值)
	Remark      string  `json:"remark" form:"remark" gorm:"comment:备注"`                          // 备注
}
