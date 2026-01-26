package k8sgpu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ImageRegistry 镜像库模型
type ImageRegistry struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"size:100;not null;comment:镜像名称"`       // 镜像名称(必填)
	Address     string `json:"address" form:"address" gorm:"size:500;not null;comment:镜像地址"`  // 镜像地址(必填)
	Description string `json:"description" form:"description" gorm:"size:500;comment:镜像描述"`   // 镜像描述
	Source      string `json:"source" form:"source" gorm:"size:100;comment:镜像来源"`            // 镜像来源
	IsPublished bool   `json:"isPublished" form:"isPublished" gorm:"default:true;comment:是否上架"` // 是否上架(默认上架)
	Remark      string `json:"remark" form:"remark" gorm:"size:500;comment:备注"`               // 备注
}

// TableName 指定表名
func (ImageRegistry) TableName() string {
	return "k8s_gpu_image_registries"
}
