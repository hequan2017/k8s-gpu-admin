package k8sgpu

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/k8sgpu"
	k8sgpuReq "github.com/flipped-aurora/gin-vue-admin/server/model/k8sgpu/request"
	"github.com/pkg/errors"
)

type InstanceService struct{}

var InstanceServiceApp = new(InstanceService)

// CreateInstance 创建实例
func (is *InstanceService) CreateInstance(req k8sgpuReq.InstanceAdd, userID uint) (instance k8sgpu.Instance, err error) {
	// 获取镜像信息
	var image k8sgpu.ImageRegistry
	err = global.GVA_DB.Where("id = ?", req.ImageID).First(&image).Error
	if err != nil {
		return instance, errors.New("镜像不存在")
	}

	// 获取规格信息
	var spec k8sgpu.ProductSpec
	err = global.GVA_DB.Where("id = ?", req.SpecID).First(&spec).Error
	if err != nil {
		return instance, errors.New("产品规格不存在")
	}

	// 匹配可用节点
	availableNodes, err := ComputeNodeServiceApp.MatchAvailableNodes(req.SpecID)
	if err != nil {
		return instance, errors.Wrap(err, "匹配节点失败")
	}
	if len(availableNodes) == 0 {
		return instance, errors.New("没有可用的算力节点")
	}
	// 选择第一个可用节点
	selectedNode := availableNodes[0]

	// 创建实例记录
	instance.InstanceName = req.InstanceName
	instance.ImageID = req.ImageID
	instance.SpecID = req.SpecID
	instance.NodeID = selectedNode.ID
	instance.UserID = userID
	instance.Remark = req.Remark
	instance.ContainerStatus = "creating"

	err = global.GVA_DB.Create(&instance).Error
	if err != nil {
		return instance, errors.Wrap(err, "创建实例记录失败")
	}

	// 创建Docker容器
	containerID, err := is.createDockerContainer(selectedNode, image.Address, instance.InstanceName, spec)
	if err != nil {
		// 创建容器失败，删除实例记录
		global.GVA_DB.Delete(&instance)
		return instance, errors.Wrap(err, "创建Docker容器失败")
	}

	// 更新实例的容器ID和状态
	instance.ContainerID = containerID
	instance.ContainerStatus = "running"
	err = global.GVA_DB.Save(&instance).Error

	return instance, err
}

// createDockerContainer 创建Docker容器
func (is *InstanceService) createDockerContainer(node k8sgpu.ComputeNode, imageURL, containerName string, spec k8sgpu.ProductSpec) (containerID string, err error) {
	// 创建Docker客户端
	cli, err := is.createDockerClient(node)
	if err != nil {
		return "", err
	}
	defer cli.Close()

	// 拉取镜像
	reader, err := cli.ImagePull(context.Background(), imageURL, types.ImagePullOptions{})
	if err != nil {
		return "", errors.Wrap(err, "拉取镜像失败")
	}
	io.Copy(io.Discard, reader)
	reader.Close()

	// 配置容器创建选项
	config := &container.Config{
		Image: imageURL,
	}

	// 添加GPU配置
	if spec.GPUCount > 0 {
		config.DeviceRequests = []container.DeviceRequest{
			{
				Driver: "nvidia",
				Count:  spec.GPUCount,
				Capabilities: [][]string{
					{"gpu"},
					{"nvidia"},
				},
			},
		}
	}

	// 配置主机和资源限制
	hostConfig := &container.HostConfig{
		RestartPolicy: container.RestartPolicy{
			Name: "unless-stopped",
		},
	}

	// CPU配置
	if spec.CPUCores > 0 {
		hostConfig.NanoCPUs = int64(spec.CPUCores * 1e9)
	}

	// 内存配置
	if spec.Memory > 0 {
		hostConfig.Memory = int64(spec.Memory * 1024 * 1024 * 1024) // GB转字节
	}

	// 系统盘配置(overlay2.size)
	// 注意：需要存储驱动支持overlay2，如果不支持则忽略
	hostConfig.StorageOpt = map[string]string{
		"size": fmt.Sprintf("%dG", spec.SystemDisk),
	}

	// 数据盘配置(命名卷映射)
	if spec.DataDisk > 0 {
		volumeName := fmt.Sprintf("instance-data-%s", containerName)
		hostConfig.Binds = []string{
			fmt.Sprintf("%s:/data", volumeName),
		}
	}

	// 创建容器
	resp, err := cli.ContainerCreate(context.Background(), config, hostConfig, nil, nil, containerName)
	if err != nil {
		// 如果storage_opt不支持，重试不带storage_opt的配置
		if strings.Contains(err.Error(), "storage-opt") || strings.Contains(err.Error(), "overlay2.size") {
			hostConfig.StorageOpt = nil
			resp, err = cli.ContainerCreate(context.Background(), config, hostConfig, nil, nil, containerName)
			if err != nil {
				return "", errors.Wrap(err, "创建容器失败")
			}
		} else {
			return "", errors.Wrap(err, "创建容器失败")
		}
	}

	// 启动容器
	err = cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	if err != nil {
		// 启动失败，删除容器
		cli.ContainerRemove(context.Background(), resp.ID, types.ContainerRemoveOptions{})
		return "", errors.Wrap(err, "启动容器失败")
	}

	return resp.ID, nil
}

// createDockerClient 创建Docker客户端
func (is *InstanceService) createDockerClient(node k8sgpu.ComputeNode) (*client.Client, error) {
	endpoint := node.DockerEndpoint
	if endpoint == "" {
		// 如果没有配置Docker连接地址，使用默认的unix socket
		endpoint = "unix:///var/run/docker.sock"
	}

	// 如果是TCP连接，使用HTTP协议
	if strings.HasPrefix(endpoint, "tcp://") || strings.HasPrefix(endpoint, "http://") || strings.HasPrefix(endpoint, "https://") {
		endpoint = strings.TrimPrefix(endpoint, "tcp://")
		endpoint = strings.TrimPrefix(endpoint, "http://")
		endpoint = strings.TrimPrefix(endpoint, "https://")
		if !strings.HasPrefix(endpoint, "http://") && !strings.HasPrefix(endpoint, "https://") {
			if node.UseTLS {
				endpoint = "https://" + endpoint
			} else {
				endpoint = "http://" + endpoint
			}
		}
	}

	httpClient := &http.Client{}
	if node.UseTLS && node.TLSCert != "" && node.TLSKey != "" {
		// TODO: 配置TLS证书
	}

	cli, err := client.NewClientWithOpts(client.WithHost(endpoint), client.WithHTTPClient(httpClient))
	if err != nil {
		return nil, errors.Wrap(err, "创建Docker客户端失败")
	}

	return cli, nil
}

// DeleteInstance 删除实例
func (is *InstanceService) DeleteInstance(id uint, userID uint, isAdmin bool) (err error) {
	var instance k8sgpu.Instance
	err = global.GVA_DB.Where("id = ?", id).First(&instance).Error
	if err != nil {
		return errors.New("实例不存在")
	}

	// 权限检查
	if !isAdmin && instance.UserID != userID {
		return errors.New("无权删除此实例")
	}

	// 如果有容器ID，先删除容器
	if instance.ContainerID != "" {
		var node k8sgpu.ComputeNode
		err = global.GVA_DB.Where("id = ?", instance.NodeID).First(&node).Error
		if err == nil {
			err = is.deleteDockerContainer(node, instance.ContainerID)
			if err != nil {
				return errors.Wrap(err, "删除Docker容器失败，请先手动删除容器")
			}
		}
	}

	// 删除实例记录
	err = global.GVA_DB.Delete(&instance).Error
	return err
}

// deleteDockerContainer 删除Docker容器
func (is *InstanceService) deleteDockerContainer(node k8sgpu.ComputeNode, containerID string) error {
	cli, err := is.createDockerClient(node)
	if err != nil {
		return err
	}
	defer cli.Close()

	// 强制删除容器及其数据卷
	err = cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{
		Force: true,
		RemoveVolumes: true,
	})
	if err != nil {
		return errors.Wrap(err, "删除容器失败")
	}

	return nil
}

// GetInstance 获取实例信息
func (is *InstanceService) GetInstance(id uint, userID uint, isAdmin bool) (instance k8sgpu.Instance, err error) {
	err = global.GVA_DB.Preload("Image").Preload("Spec").Preload("Node").Preload("User").Where("id = ?", id).First(&instance).Error
	if err != nil {
		return instance, errors.New("实例不存在")
	}

	// 权限检查
	if !isAdmin && instance.UserID != userID {
		return instance, errors.New("无权查看此实例")
	}

	return instance, nil
}

// GetInstanceList 分页获取实例列表
func (is *InstanceService) GetInstanceList(info k8sgpuReq.InstanceSearch, userID uint, isAdmin bool) (list []k8sgpu.Instance, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&k8sgpu.Instance{})

	// 权限过滤：非管理员只能查看自己的实例
	if !isAdmin {
		db = db.Where("user_id = ?", userID)
	}

	if info.InstanceName != "" {
		db = db.Where("instance_name LIKE ?", "%"+info.InstanceName+"%")
	}
	if info.ContainerStatus != "" {
		db = db.Where("container_status = ?", info.ContainerStatus)
	}
	if info.ImageID > 0 {
		db = db.Where("image_id = ?", info.ImageID)
	}
	if info.SpecID > 0 {
		db = db.Where("spec_id = ?", info.SpecID)
	}
	if info.NodeID > 0 {
		db = db.Where("node_id = ?", info.NodeID)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Preload("Image").Preload("Spec").Preload("Node").Preload("User").
		Limit(limit).Offset(offset).Order("id DESC").Find(&list).Error
	return list, total, err
}

// ContainerAction 容器操作
func (is *InstanceService) ContainerAction(id uint, action string, userID uint, isAdmin bool) (err error) {
	var instance k8sgpu.Instance
	err = global.GVA_DB.Where("id = ?", id).First(&instance).Error
	if err != nil {
		return errors.New("实例不存在")
	}

	// 权限检查
	if !isAdmin && instance.UserID != userID {
		return errors.New("无权操作此实例")
	}

	if instance.ContainerID == "" {
		return errors.New("容器ID不存在")
	}

	var node k8sgpu.ComputeNode
	err = global.GVA_DB.Where("id = ?", instance.NodeID).First(&node).Error
	if err != nil {
		return errors.New("节点不存在")
	}

	cli, err := is.createDockerClient(node)
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx := context.Background()

	switch action {
	case "start":
		err = cli.ContainerStart(ctx, instance.ContainerID, types.ContainerStartOptions{})
		if err != nil {
			return errors.Wrap(err, "启动容器失败")
		}
		instance.ContainerStatus = "running"
	case "stop":
		err = cli.ContainerStop(ctx, instance.ContainerID, container.StopOptions{})
		if err != nil {
			return errors.Wrap(err, "停止容器失败")
		}
		instance.ContainerStatus = "stopped"
	case "restart":
		err = cli.ContainerRestart(ctx, instance.ContainerID, container.StopOptions{})
		if err != nil {
			return errors.Wrap(err, "重启容器失败")
		}
		instance.ContainerStatus = "running"
	default:
		return errors.New("不支持的操作")
	}

	global.GVA_DB.Save(&instance)
	return nil
}

// GetContainerLogs 获取容器日志
func (is *InstanceService) GetContainerLogs(id uint, userID uint, isAdmin bool) (logs string, err error) {
	var instance k8sgpu.Instance
	err = global.GVA_DB.Where("id = ?", id).First(&instance).Error
	if err != nil {
		return "", errors.New("实例不存在")
	}

	// 权限检查
	if !isAdmin && instance.UserID != userID {
		return "", errors.New("无权查看此实例")
	}

	if instance.ContainerID == "" {
		return "", errors.New("容器ID不存在")
	}

	var node k8sgpu.ComputeNode
	err = global.GVA_DB.Where("id = ?", instance.NodeID).First(&node).Error
	if err != nil {
		return "", errors.New("节点不存在")
	}

	cli, err := is.createDockerClient(node)
	if err != nil {
		return "", err
	}
	defer cli.Close()

	reader, err := cli.ContainerLogs(context.Background(), instance.ContainerID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Tail:       "100",
	})
	if err != nil {
		return "", errors.Wrap(err, "获取容器日志失败")
	}
	defer reader.Close()

	logBytes, err := io.ReadAll(reader)
	if err != nil {
		return "", errors.Wrap(err, "读取日志失败")
	}

	return string(logBytes), nil
}

// UpdateInstanceStatus 更新实例状态
func (is *InstanceService) UpdateInstanceStatus(instanceID uint) error {
	var instance k8sgpu.Instance
	err := global.GVA_DB.Where("id = ?", instanceID).First(&instance).Error
	if err != nil {
		return err
	}

	if instance.ContainerID == "" {
		return nil
	}

	var node k8sgpu.ComputeNode
	err = global.GVA_DB.Where("id = ?", instance.NodeID).First(&node).Error
	if err != nil {
		return err
	}

	cli, err := is.createDockerClient(node)
	if err != nil {
		return err
	}
	defer cli.Close()

	containerJSON, err := cli.ContainerInspect(context.Background(), instance.ContainerID)
	if err != nil {
		instance.ContainerStatus = "removed"
	} else {
		if containerJSON.State.Running {
			instance.ContainerStatus = "running"
		} else {
			instance.ContainerStatus = "stopped"
		}
	}

	global.GVA_DB.Save(&instance)
	return nil
}
