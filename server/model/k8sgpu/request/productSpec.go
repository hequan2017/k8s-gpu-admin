package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// ProductSpecSearch 产品规格搜索结构体
type ProductSpecSearch struct {
	Name        string `json:"name" form:"name" gorm:"like;comment:规格名称"`        // 规格名称(模糊查询)
	GPUModel    string `json:"gpuModel" form:"gpuModel" gorm:"comment:显卡型号"`     // 显卡型号
	IsPublished *bool  `json:"isPublished" form:"isPublished" gorm:"comment:是否上架"` // 是否上架
	request.PageInfo
}

// ProductSpecAdd 产品规格添加请求
type ProductSpecAdd struct {
	Name         string  `json:"name" binding:"required" form:"name" gorm:"comment:规格名称"`                    // 规格名称(必填)
	GPUModel     string  `json:"gpuModel" binding:"required" form:"gpuModel" gorm:"comment:显卡型号"`            // 显卡型号(必填)
	GPUCount     int     `json:"gpuCount" form:"gpuCount" gorm:"default:1;comment:显卡数量"`                    // 显卡数量
	CPUCores     int     `json:"cpuCores" form:"cpuCores" gorm:"default:0;comment:CPU核心数"`                  // CPU核心数
	Memory       int     `json:"memory" form:"memory" gorm:"default:0;comment:内存(GB)"`                       // 内存(GB)
	SystemDisk   int     `json:"systemDisk" form:"systemDisk" gorm:"default:0;comment:系统盘容量(GB)"`           // 系统盘容量(GB)
	DataDisk     int     `json:"dataDisk" form:"dataDisk" gorm:"default:0;comment:数据盘容量(GB)"`               // 数据盘容量(GB)
	PricePerHour float64 `json:"pricePerHour" form:"pricePerHour" gorm:"type:decimal(10,4);default:0;comment:价格/小时"` // 价格/小时
	IsPublished  bool    `json:"isPublished" form:"isPublished" gorm:"default:true;comment:是否上架"`           // 是否上架(默认上架)
	Remark       string  `json:"remark" form:"remark" gorm:"comment:备注"`                                      // 备注
}

// ProductSpecUpdate 产品规格更新请求
type ProductSpecUpdate struct {
	ID           uint     `json:"id" binding:"required" form:"id" gorm:"primarykey;comment:主键ID"`     // 主键ID
	Name         string   `json:"name" binding:"required" form:"name" gorm:"comment:规格名称"`          // 规格名称(必填)
	GPUModel     string   `json:"gpuModel" binding:"required" form:"gpuModel" gorm:"comment:显卡型号"`  // 显卡型号(必填)
	GPUCount     int      `json:"gpuCount" form:"gpuCount" gorm:"default:1;comment:显卡数量"`          // 显卡数量
	CPUCores     int      `json:"cpuCores" form:"cpuCores" gorm:"default:0;comment:CPU核心数"`         // CPU核心数
	Memory       int      `json:"memory" form:"memory" gorm:"default:0;comment:内存(GB)"`              // 内存(GB)
	SystemDisk   int      `json:"systemDisk" form:"systemDisk" gorm:"default:0;comment:系统盘容量(GB)"`  // 系统盘容量(GB)
	DataDisk     int      `json:"dataDisk" form:"dataDisk" gorm:"default:0;comment:数据盘容量(GB)"`      // 数据盘容量(GB)
	PricePerHour float64  `json:"pricePerHour" form:"pricePerHour" gorm:"type:decimal(10,4);default:0;comment:价格/小时"` // 价格/小时
	IsPublished  *bool    `json:"isPublished" form:"isPublished" gorm:"comment:是否上架"`                // 是否上架(使用指针以区分零值)
	Remark       string   `json:"remark" form:"remark" gorm:"comment:备注"`                             // 备注
}
