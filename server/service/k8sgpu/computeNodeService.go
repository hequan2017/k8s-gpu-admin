package k8sgpu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/k8sgpu"
	k8sgpuReq "github.com/flipped-aurora/gin-vue-admin/server/model/k8sgpu/request"
	"github.com/pkg/errors"
)

type ComputeNodeService struct{}

var ComputeNodeServiceApp = new(ComputeNodeService)

// CreateComputeNode 创建算力节点
func (cns *ComputeNodeService) CreateComputeNode(req k8sgpuReq.ComputeNodeAdd) (err error) {
	var node k8sgpu.ComputeNode
	node.Name = req.Name
	node.Region = req.Region
	node.CPU = req.CPU
	node.Memory = req.Memory
	node.SystemDisk = req.SystemDisk
	node.DataDisk = req.DataDisk
	node.PublicIP = req.PublicIP
	node.PrivateIP = req.PrivateIP
	node.SSHPort = req.SSHPort
	node.Username = req.Username
	node.Password = req.Password
	node.GPUName = req.GPUName
	node.GPUCount = req.GPUCount
	node.DockerEndpoint = req.DockerEndpoint
	node.UseTLS = req.UseTLS
	node.TLSCACert = req.TLSCACert
	node.TLSCert = req.TLSCert
	node.TLSKey = req.TLSKey
	node.IsPublished = req.IsPublished
	node.Remark = req.Remark
	err = global.GVA_DB.Create(&node).Error
	return err
}

// DeleteComputeNode 删除算力节点
func (cns *ComputeNodeService) DeleteComputeNode(id uint) (err error) {
	// 检查是否有关联的实例
	var count int64
	global.GVA_DB.Model(&k8sgpu.Instance{}).Where("node_id = ?", id).Count(&count)
	if count > 0 {
		return errors.New("该节点下还有实例，无法删除")
	}
	err = global.GVA_DB.Delete(&k8sgpu.ComputeNode{}, id).Error
	return err
}

// UpdateComputeNode 更新算力节点
func (cns *ComputeNodeService) UpdateComputeNode(req k8sgpuReq.ComputeNodeUpdate) (err error) {
	var node k8sgpu.ComputeNode
	node.ID = req.ID
	node.Name = req.Name
	node.Region = req.Region
	node.CPU = req.CPU
	node.Memory = req.Memory
	node.SystemDisk = req.SystemDisk
	node.DataDisk = req.DataDisk
	node.PublicIP = req.PublicIP
	node.PrivateIP = req.PrivateIP
	node.SSHPort = req.SSHPort
	node.Username = req.Username
	node.Password = req.Password
	node.GPUName = req.GPUName
	node.GPUCount = req.GPUCount
	node.DockerEndpoint = req.DockerEndpoint
	if req.UseTLS != nil {
		node.UseTLS = *req.UseTLS
	}
	node.TLSCACert = req.TLSCACert
	node.TLSCert = req.TLSCert
	node.TLSKey = req.TLSKey
	if req.IsPublished != nil {
		node.IsPublished = *req.IsPublished
	}
	node.Remark = req.Remark
	err = global.GVA_DB.Save(&node).Error
	return err
}

// GetComputeNode 获取算力节点信息
func (cns *ComputeNodeService) GetComputeNode(id uint) (node k8sgpu.ComputeNode, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&node).Error
	return
}

// GetComputeNodeList 分页获取算力节点列表
func (cns *ComputeNodeService) GetComputeNodeList(info k8sgpuReq.ComputeNodeSearch) (list []k8sgpu.ComputeNode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&k8sgpu.ComputeNode{})

	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Region != "" {
		db = db.Where("region = ?", info.Region)
	}
	if info.GPUName != "" {
		db = db.Where("gpu_name = ?", info.GPUName)
	}
	if info.IsPublished != nil {
		db = db.Where("is_published = ?", *info.IsPublished)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id DESC").Find(&list).Error
	return list, total, err
}

// MatchAvailableNodes 匹配可用节点(根据产品规格)
func (cns *ComputeNodeService) MatchAvailableNodes(specID uint) (nodes []k8sgpu.ComputeNode, err error) {
	// 获取产品规格
	var spec k8sgpu.ProductSpec
	err = global.GVA_DB.Where("id = ?", specID).First(&spec).Error
	if err != nil {
		return
	}

	// 获取所有已上架的节点
	var allNodes []k8sgpu.ComputeNode
	err = global.GVA_DB.Where("is_published = ?", true).Find(&allNodes).Error
	if err != nil {
		return
	}

	// 获取所有已创建的实例
	var instances []k8sgpu.Instance
	err = global.GVA_DB.Preload("Spec").Find(&instances).Error
	if err != nil {
		return
	}

	// 计算每个节点的已用资源
	nodeUsedMap := make(map[uint]ResourceUsed)
	for _, inst := range instances {
		if inst.NodeID == 0 || inst.SpecID == 0 {
			continue
		}
		used, ok := nodeUsedMap[inst.NodeID]
		if !ok {
			used = ResourceUsed{}
		}
		// 累加已用资源
		used.CPU += inst.Spec.CPUCores
		used.Memory += inst.Spec.Memory
		used.SystemDisk += inst.Spec.SystemDisk
		used.DataDisk += inst.Spec.DataDisk
		used.GPUCount += inst.Spec.GPUCount
		nodeUsedMap[inst.NodeID] = used
	}

	// 筛选满足条件的节点
	for _, node := range allNodes {
		used := nodeUsedMap[node.ID]
		// 检查是否满足规格要求
		if node.CPU-used.CPU >= spec.CPUCores &&
			node.Memory-used.Memory >= spec.Memory &&
			node.SystemDisk-used.SystemDisk >= spec.SystemDisk &&
			node.DataDisk-used.DataDisk >= spec.DataDisk &&
			node.GPUCount-used.GPUCount >= spec.GPUCount &&
			node.GPUName == spec.GPUModel {
			nodes = append(nodes, node)
		}
	}

	return
}

// ResourceUsed 节点已用资源
type ResourceUsed struct {
	CPU        int
	Memory     int
	SystemDisk int
	DataDisk   int
	GPUCount   int
}
