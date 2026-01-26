package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// ComputeNodeSearch 算力节点搜索结构体
type ComputeNodeSearch struct {
	Name        string `json:"name" form:"name" gorm:"like;comment:节点名称"`        // 节点名称(模糊查询)
	Region      string `json:"region" form:"region" gorm:"comment:区域"`             // 区域
	GPUName     string `json:"gpuName" form:"gpuName" gorm:"comment:显卡名称"`        // 显卡名称
	IsPublished *bool  `json:"isPublished" form:"isPublished" gorm:"comment:是否上架"` // 是否上架
	request.PageInfo
}

// ComputeNodeAdd 算力节点添加请求
type ComputeNodeAdd struct {
	Name           string `json:"name" binding:"required" form:"name" gorm:"comment:节点名称"`                      // 节点名称(必填)
	Region         string `json:"region" form:"region" gorm:"comment:区域"`                                      // 区域
	CPU            int    `json:"cpu" form:"cpu" gorm:"default:0;comment:CPU核心数"`                             // CPU核心数
	Memory         int    `json:"memory" form:"memory" gorm:"default:0;comment:内存(GB)"`                        // 内存(GB)
	SystemDisk     int    `json:"systemDisk" form:"systemDisk" gorm:"default:0;comment:系统盘容量(GB)"`           // 系统盘容量(GB)
	DataDisk       int    `json:"dataDisk" form:"dataDisk" gorm:"default:0;comment:数据盘容量(GB)"`               // 数据盘容量(GB)
	PublicIP       string `json:"publicIP" binding:"required" form:"publicIP" gorm:"comment:公网IP地址"`          // 公网IP地址(必填)
	PrivateIP      string `json:"privateIP" binding:"required" form:"privateIP" gorm:"comment:内网IP地址"`        // 内网IP地址(必填)
	SSHPort        int    `json:"sshPort" form:"sshPort" gorm:"default:22;comment:SSH端口"`                     // SSH端口(默认22)
	Username       string `json:"username" form:"username" gorm:"comment:SSH用户名"`                            // SSH用户名
	Password       string `json:"password" form:"password" gorm:"comment:SSH密码"`                              // SSH密码
	GPUName        string `json:"gpuName" form:"gpuName" gorm:"comment:显卡名称"`                                // 显卡名称
	GPUCount       int    `json:"gpuCount" form:"gpuCount" gorm:"default:0;comment:显卡数量"`                   // 显卡数量
	DockerEndpoint string `json:"dockerEndpoint" form:"dockerEndpoint" gorm:"comment:Docker连接地址"`           // Docker连接地址
	UseTLS         bool   `json:"useTLS" form:"useTLS" gorm:"default:true;comment:是否使用TLS"`                 // 是否使用TLS(默认启用)
	TLSCACert      string `json:"tlsCaCert" form:"tlsCaCert" gorm:"type:text;comment:CA证书"`                  // CA证书
	TLSCert        string `json:"tlsCert" form:"tlsCert" gorm:"type:text;comment:客户端证书"`                  // 客户端证书
	TLSKey         string `json:"tlsKey" form:"tlsKey" gorm:"type:text;comment:客户端私钥"`                     // 客户端私钥
	IsPublished    bool   `json:"isPublished" form:"isPublished" gorm:"default:true;comment:是否上架"`          // 是否上架(默认上架)
	Remark         string `json:"remark" form:"remark" gorm:"comment:备注"`                                    // 备注
}

// ComputeNodeUpdate 算力节点更新请求
type ComputeNodeUpdate struct {
	ID             uint    `json:"id" binding:"required" form:"id" gorm:"primarykey;comment:主键ID"`    // 主键ID
	Name           string  `json:"name" binding:"required" form:"name" gorm:"comment:节点名称"`         // 节点名称(必填)
	Region         string  `json:"region" form:"region" gorm:"comment:区域"`                           // 区域
	CPU            int     `json:"cpu" form:"cpu" gorm:"default:0;comment:CPU核心数"`                  // CPU核心数
	Memory         int     `json:"memory" form:"memory" gorm:"default:0;comment:内存(GB)"`             // 内存(GB)
	SystemDisk     int     `json:"systemDisk" form:"systemDisk" gorm:"default:0;comment:系统盘容量(GB)"` // 系统盘容量(GB)
	DataDisk       int     `json:"dataDisk" form:"dataDisk" gorm:"default:0;comment:数据盘容量(GB)"`     // 数据盘容量(GB)
	PublicIP       string  `json:"publicIP" binding:"required" form:"publicIP" gorm:"comment:公网IP地址"` // 公网IP地址(必填)
	PrivateIP      string  `json:"privateIP" binding:"required" form:"privateIP" gorm:"comment:内网IP地址"` // 内网IP地址(必填)
	SSHPort        int     `json:"sshPort" form:"sshPort" gorm:"default:22;comment:SSH端口"`            // SSH端口(默认22)
	Username       string  `json:"username" form:"username" gorm:"comment:SSH用户名"`                   // SSH用户名
	Password       string  `json:"password" form:"password" gorm:"comment:SSH密码"`                     // SSH密码
	GPUName        string  `json:"gpuName" form:"gpuName" gorm:"comment:显卡名称"`                      // 显卡名称
	GPUCount       int     `json:"gpuCount" form:"gpuCount" gorm:"default:0;comment:显卡数量"`          // 显卡数量
	DockerEndpoint string  `json:"dockerEndpoint" form:"dockerEndpoint" gorm:"comment:Docker连接地址"`  // Docker连接地址
	UseTLS         *bool   `json:"useTLS" form:"useTLS" gorm:"comment:是否使用TLS"`                     // 是否使用TLS(使用指针以区分零值)
	TLSCACert      string  `json:"tlsCaCert" form:"tlsCaCert" gorm:"type:text;comment:CA证书"`         // CA证书
	TLSCert        string  `json:"tlsCert" form:"tlsCert" gorm:"type:text;comment:客户端证书"`         // 客户端证书
	TLSKey         string  `json:"tlsKey" form:"tlsKey" gorm:"type:text;comment:客户端私钥"`            // 客户端私钥
	IsPublished    *bool   `json:"isPublished" form:"isPublished" gorm:"comment:是否上架"`               // 是否上架(使用指针以区分零值)
	Remark         string  `json:"remark" form:"remark" gorm:"comment:备注"`                           // 备注
}
