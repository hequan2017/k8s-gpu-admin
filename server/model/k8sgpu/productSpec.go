package k8sgpu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ProductSpec 产品规格模型
type ProductSpec struct {
	global.GVA_MODEL
	Name          string  `json:"name" form:"name" gorm:"size:100;not null;comment:规格名称"`               // 规格名称(必填)
	GPUModel      string  `json:"gpuModel" form:"gpuModel" gorm:"size:100;not null;comment:显卡型号"`        // 显卡型号(必填)
	GPUCount      int     `json:"gpuCount" form:"gpuCount" gorm:"default:1;comment:显卡数量"`               // 显卡数量
	CPUCores      int     `json:"cpuCores" form:"cpuCores" gorm:"default:0;comment:CPU核心数"`             // CPU核心数
	Memory        int     `json:"memory" form:"memory" gorm:"default:0;comment:内存(GB)"`                  // 内存(GB)
	SystemDisk    int     `json:"systemDisk" form:"systemDisk" gorm:"default:0;comment:系统盘容量(GB)"`     // 系统盘容量(GB)
	DataDisk      int     `json:"dataDisk" form:"dataDisk" gorm:"default:0;comment:数据盘容量(GB)"`         // 数据盘容量(GB)
	PricePerHour  float64 `json:"pricePerHour" form:"pricePerHour" gorm:"type:decimal(10,4);default:0;comment:价格/小时"` // 价格/小时
	IsPublished   bool    `json:"isPublished" form:"isPublished" gorm:"default:true;comment:是否上架"`      // 是否上架(默认上架)
	Remark        string  `json:"remark" form:"remark" gorm:"size:500;comment:备注"`                       // 备注
}

// TableName 指定表名
func (ProductSpec) TableName() string {
	return "k8s_gpu_product_specs"
}
