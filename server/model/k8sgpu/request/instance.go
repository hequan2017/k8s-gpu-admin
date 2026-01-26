package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// InstanceSearch 实例搜索结构体
type InstanceSearch struct {
	InstanceName     string `json:"instanceName" form:"instanceName" gorm:"like;comment:实例名称"`        // 实例名称(模糊查询)
	ContainerStatus  string `json:"containerStatus" form:"containerStatus" gorm:"comment:容器状态"`        // 容器状态
	ImageID          uint   `json:"imageId" form:"imageId" gorm:"comment:镜像ID"`                        // 镜像ID
	SpecID           uint   `json:"specId" form:"specId" gorm:"comment:产品规格ID"`                      // 产品规格ID
	NodeID           uint   `json:"nodeId" form:"nodeId" gorm:"comment:算力节点ID"`                      // 算力节点ID
	request.PageInfo
}

// InstanceAdd 实例添加请求
type InstanceAdd struct {
	InstanceName string `json:"instanceName" binding:"required" form:"instanceName" gorm:"comment:实例名称"` // 实例名称(必填)
	ImageID      uint   `json:"imageId" binding:"required" form:"imageId" gorm:"comment:镜像ID"`           // 镜像ID(必填)
	SpecID       uint   `json:"specId" binding:"required" form:"specId" gorm:"comment:产品规格ID"`         // 产品规格ID(必填)
	Remark       string `json:"remark" form:"remark" gorm:"comment:备注"`                                 // 备注
}

// InstanceUpdate 实例更新请求
type InstanceUpdate struct {
	ID           uint   `json:"id" binding:"required" form:"id" gorm:"primarykey;comment:主键ID"`    // 主键ID
	InstanceName string `json:"instanceName" binding:"required" form:"instanceName" gorm:"comment:实例名称"` // 实例名称(必填)
	ImageID      uint   `json:"imageId" binding:"required" form:"imageId" gorm:"comment:镜像ID"`    // 镜像ID(必填)
	SpecID       uint   `json:"specId" binding:"required" form:"specId" gorm:"comment:产品规格ID"`  // 产品规格ID(必填)
	Remark       string `json:"remark" form:"remark" gorm:"comment:备注"`                          // 备注
}

// MatchNodeRequest 匹配节点请求
type MatchNodeRequest struct {
	SpecID uint `json:"specId" binding:"required" form:"specId" gorm:"comment:产品规格ID"` // 产品规格ID(必填)
}

// ContainerActionRequest 容器操作请求
type ContainerActionRequest struct {
	ID     uint   `json:"id" binding:"required" form:"id" gorm:"comment:实例ID"` // 实例ID(必填)
	Action string `json:"action" binding:"required" form:"action" gorm:"comment:操作类型"` // 操作类型: start/stop/restart/logs/exec
}
