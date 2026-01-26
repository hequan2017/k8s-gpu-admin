package k8sgpu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// Instance 实例管理模型
type Instance struct {
	global.GVA_MODEL
	InstanceName   string         `json:"instanceName" form:"instanceName" gorm:"size:100;not null;comment:实例名称"`  // 实例名称(必填)
	ImageID        uint           `json:"imageId" form:"imageId" gorm:"not null;comment:镜像ID"`                      // 镜像ID(必填)
	Image          ImageRegistry  `json:"image" gorm:"foreignKey:ImageID;comment:镜像详情"`                           // 镜像详情
	SpecID         uint           `json:"specId" form:"specId" gorm:"not null;comment:产品规格ID"`                    // 产品规格ID(必填)
	Spec           ProductSpec    `json:"spec" gorm:"foreignKey:SpecID;comment:规格详情"`                              // 规格详情
	NodeID         uint           `json:"nodeId" form:"nodeId" gorm:"not null;comment:算力节点ID"`                     // 算力节点ID(必填)
	Node           ComputeNode    `json:"node" gorm:"foreignKey:NodeID;comment:节点详情"`                             // 节点详情
	UserID         uint           `json:"userId" form:"userId" gorm:"not null;comment:用户ID"`                       // 用户ID(后端自动填写)
	User           system.SysUser `json:"user" gorm:"foreignKey:UserID;comment:用户详情"`                             // 用户详情
	ContainerID    string         `json:"containerId" form:"containerId" gorm:"size:200;comment:容器ID"`              // Docker容器ID(后端自动回填)
	ContainerStatus string        `json:"containerStatus" form:"containerStatus" gorm:"size:50;comment:容器状态"`     // 容器状态(后端自动填写)
	Remark         string         `json:"remark" form:"remark" gorm:"size:500;comment:备注"`                          // 备注
}

// TableName 指定表名
func (Instance) TableName() string {
	return "k8s_gpu_instances"
}
