package k8sgpu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ComputeNode 算力节点模型
type ComputeNode struct {
	global.GVA_MODEL
	Name             string `json:"name" form:"name" gorm:"size:100;not null;comment:节点名称"`                              // 节点名称(必填)
	Region           string `json:"region" form:"region" gorm:"size:100;comment:区域"`                                     // 区域
	CPU              int    `json:"cpu" form:"cpu" gorm:"default:0;comment:CPU核心数"`                                     // CPU核心数
	Memory           int    `json:"memory" form:"memory" gorm:"default:0;comment:内存(GB)"`                                // 内存(GB)
	SystemDisk       int    `json:"systemDisk" form:"systemDisk" gorm:"default:0;comment:系统盘容量(GB)"`                    // 系统盘容量(GB)
	DataDisk         int    `json:"dataDisk" form:"dataDisk" gorm:"default:0;comment:数据盘容量(GB)"`                        // 数据盘容量(GB)
	PublicIP         string `json:"publicIP" form:"publicIP" gorm:"size:50;not null;comment:公网IP地址"`                     // 公网IP地址(必填)
	PrivateIP        string `json:"privateIP" form:"privateIP" gorm:"size:50;not null;comment:内网IP地址"`                   // 内网IP地址(必填)
	SSHPort          int    `json:"sshPort" form:"sshPort" gorm:"default:22;comment:SSH端口"`                              // SSH端口(默认22)
	Username         string `json:"username" form:"username" gorm:"size:50;comment:SSH用户名"`                             // SSH用户名
	Password         string `json:"password" form:"password" gorm:"size:200;comment:SSH密码"`                               // SSH密码
	GPUName          string `json:"gpuName" form:"gpuName" gorm:"size:100;comment:显卡名称"`                                // 显卡名称
	GPUCount         int    `json:"gpuCount" form:"gpuCount" gorm:"default:0;comment:显卡数量"`                              // 显卡数量
	DockerEndpoint   string `json:"dockerEndpoint" form:"dockerEndpoint" gorm:"size:200;comment:Docker连接地址"`            // Docker连接地址
	UseTLS           bool   `json:"useTLS" form:"useTLS" gorm:"default:true;comment:是否使用TLS"`                           // 是否使用TLS(默认启用)
	TLSCACert        string `json:"tlsCaCert" form:"tlsCaCert" type:"text" gorm:"type:text;comment:CA证书"`                // CA证书
	TLSCert          string `json:"tlsCert" form:"tlsCert" type:"text" gorm:"type:text;comment:客户端证书"`                 // 客户端证书
	TLSKey           string `json:"tlsKey" form:"tlsKey" type:"text" gorm:"type:text;comment:客户端私钥"`                   // 客户端私钥
	IsPublished      bool   `json:"isPublished" form:"isPublished" gorm:"default:true;comment:是否上架"`                    // 是否上架(默认上架)
	Remark           string `json:"remark" form:"remark" gorm:"size:500;comment:备注"`                                     // 备注
}

// TableName 指定表名
func (ComputeNode) TableName() string {
	return "k8s_gpu_compute_nodes"
}
