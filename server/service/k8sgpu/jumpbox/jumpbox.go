package jumpbox

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/k8sgpu"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

// JumpboxService SSH跳板机服务
type JumpboxService struct {
	Listener net.Listener
	Config   *ssh.ServerConfig
	Running  bool
}

var JumpboxServiceApp = new(JumpboxService)

// Start 启动SSH跳板机服务器
func (js *JumpboxService) Start(port int) error {
	// 配置SSH服务器
	config := &ssh.ServerConfig{
		PasswordCallback: js.passwordCallback,
		// 可以添加公钥认证
		// PublicKeyCallback: ...
	}

	// 生成主机密钥
	key, err := js.generateHostKey()
	if err != nil {
		return fmt.Errorf("生成主机密钥失败: %w", err)
	}
	config.AddHostKey(key)

	js.Config = config

	// 监听端口
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("监听端口失败: %w", err)
	}
	js.Listener = listener
	js.Running = true

	global.GVA_LOG.Info(fmt.Sprintf("SSH跳板机服务已启动，监听端口: %d", port))

	// 接受连接
	go js.acceptConnections()

	return nil
}

// Stop 停止SSH跳板机服务器
func (js *JumpboxService) Stop() error {
	js.Running = false
	if js.Listener != nil {
		return js.Listener.Close()
	}
	return nil
}

// acceptConnections 接受客户端连接
func (js *JumpboxService) acceptConnections() {
	for js.Running {
		conn, err := js.Listener.Accept()
		if err != nil {
			if js.Running {
				global.GVA_LOG.Error("接受连接失败", zap.Error(err))
			}
			continue
		}

		go js.handleConnection(conn)
	}
}

// passwordCallback 密码认证回调
func (js *JumpboxService) passwordCallback(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
	username := conn.User()

	// 查询用户
	var user system.SysUser
	err := global.GVA_DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		global.GVA_LOG.Warn(fmt.Sprintf("用户 %s 认证失败: 用户不存在", username))
		return nil, fmt.Errorf("用户不存在")
	}

	// 验证密码
	if ok := utils.BcryptCheck(string(password), user.Password); !ok {
		global.GVA_LOG.Warn(fmt.Sprintf("用户 %s 认证失败: 密码错误", username))
		return nil, fmt.Errorf("密码错误")
	}

	global.GVA_LOG.Info(fmt.Sprintf("用户 %s 认证成功", username))
	return nil, nil
}

// handleConnection 处理SSH连接
func (js *JumpboxService) handleConnection(conn net.Conn) {
	defer conn.Close()

	// 建立SSH连接
	sshConn, chans, reqs, err := ssh.NewServerConn(conn, js.Config)
	if err != nil {
		global.GVA_LOG.Error("建立SSH连接失败", zap.Error(err))
		return
	}
	defer sshConn.Close()

	username := sshConn.User()
	global.GVA_LOG.Info(fmt.Sprintf("用户 %s 已连接", username))

	// 处理全局请求
	go ssh.DiscardRequests(reqs)

	// 处理通道请求
	for newChannel := range chans {
		if newChannel.ChannelType() != "session" {
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}

		channel, requests, err := newChannel.Accept()
		if err != nil {
			global.GVA_LOG.Error("接受通道失败", zap.Error(err))
			continue
		}

		go js.handleSession(channel, requests, username)
	}
}

// handleSession 处理会话
func (js *JumpboxService) handleSession(channel ssh.Channel, in <-chan *ssh.Request, username string) {
	defer channel.Close()

	// 获取用户信息
	var user system.SysUser
	err := global.GVA_DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		channel.Write([]byte("获取用户信息失败\r\n"))
		return
	}

	isAdmin := user.AuthorityId == 888 // 假设888是管理员权限ID

	for req := range in {
		switch req.Type {
		case "shell":
			// 拒绝shell请求
			req.Reply(true, nil)
		case "pty-req":
			// 接受终端请求
			req.Reply(true, nil)
		case "exec":
			// 处理执行请求
			go js.handleInteractive(channel, user.ID, isAdmin)
			req.Reply(true, nil)
		default:
			req.Reply(false, nil)
		}
	}
}

// handleInteractive 处理交互式会话
func (js *JumpboxService) handleInteractive(channel ssh.Channel, userID uint, isAdmin bool) {
	// 显示欢迎信息
	channel.Write([]byte("欢迎使用GPU算力管理平台 SSH跳板机\r\n\r\n"))

	for {
		// 获取用户可操作的实例列表
		var instances []k8sgpu.Instance
		query := global.GVA_DB.Model(&k8sgpu.Instance{}).Preload("Node")

		if !isAdmin {
			query = query.Where("user_id = ?", userID)
		}

		err := query.Order("id DESC").Find(&instances).Error
		if err != nil {
			channel.Write([]byte(fmt.Sprintf("获取实例列表失败: %s\r\n", err.Error())))
			return
		}

		if len(instances) == 0 {
			channel.Write([]byte("您暂无可操作的实例\r\n"))
			return
		}

		// 显示实例列表
		channel.Write([]byte("可操作的实例列表:\r\n"))
		channel.Write([]byte("序号\t实例名称\t容器ID\t算力节点\t状态\r\n"))
		channel.Write([]byte("----\r\n"))

		for i, inst := range instances {
			status := inst.ContainerStatus
			if status == "" {
				status = "unknown"
			}
			nodeName := "未知节点"
			if inst.Node.ID != 0 {
				nodeName = inst.Node.Name
			}
			channel.Write([]byte(fmt.Sprintf("%d\t%s\t%s\t%s\t%s\r\n",
				i+1, inst.InstanceName, inst.ContainerID, nodeName, status)))
		}
		channel.Write([]byte("\r\n"))

		// 提示用户选择
		channel.Write([]byte("请输入实例序号进行连接 (输入0退出): "))

		// 读取用户输入
		reader := bufio.NewReader(channel)
		line, err := reader.ReadString('\n')
		if err != nil {
			global.GVA_LOG.Error("读取用户输入失败", zap.Error(err))
			return
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// 处理退出
		if line == "0" || line == "exit" || line == "quit" {
			channel.Write([]byte("再见!\r\n"))
			return
		}

		// 解析序号
		index, err := strconv.Atoi(line)
		if err != nil {
			channel.Write([]byte(fmt.Sprintf("无效的输入: %s\r\n\r\n", line)))
			continue
		}

		if index < 1 || index > len(instances) {
			channel.Write([]byte(fmt.Sprintf("序号超出范围: %d\r\n\r\n", index)))
			continue
		}

		// 连接到容器
		selectedInstance := instances[index-1]
		err = js.connectToContainer(channel, selectedInstance)
		if err != nil {
			channel.Write([]byte(fmt.Sprintf("连接容器失败: %s\r\n\r\n", err.Error())))
			continue
		}
	}
}

// connectToContainer 连接到容器
func (js *JumpboxService) connectToContainer(channel ssh.Channel, instance k8sgpu.Instance) error {
	if instance.ContainerID == "" {
		return fmt.Errorf("容器ID为空")
	}

	// 获取节点信息
	var node k8sgpu.ComputeNode
	err := global.GVA_DB.Where("id = ?", instance.NodeID).First(&node).Error
	if err != nil {
		return fmt.Errorf("获取节点信息失败: %w", err)
	}

	channel.Write([]byte(fmt.Sprintf("\r\n正在连接到容器 %s...\r\n", instance.ContainerID)))
	channel.Write([]byte(fmt.Sprintf("节点: %s\r\n", node.Name)))
	channel.Write([]byte("\r\n"))

	// TODO: 实现Docker Exec连接到容器
	// 这里需要通过Docker API创建一个exec实例，然后转发SSH会话的输入输出

	// 临时显示消息
	channel.Write([]byte("容器终端功能正在开发中...\r\n"))
	channel.Write([]byte("容器ID: " + instance.ContainerID + "\r\n"))
	channel.Write([]byte("\r\n按Enter键返回...\r\n"))

	// 等待用户按键
	reader := bufio.NewReader(channel)
	reader.ReadString('\n')

	return nil
}

// generateHostKey 生成或加载主机密钥
func (js *JumpboxService) generateHostKey() (ssh.Signer, error) {
	keyFile := "jumpbox_host_key"

	// 尝试读取现有的密钥文件
	key, err := os.ReadFile(keyFile)
	if err == nil {
		signer, err := ssh.ParsePrivateKey(key)
		if err == nil {
			return signer, nil
		}
	}

	// 生成新的密钥
	// 这里简化处理，实际应用中应该使用更安全的方式生成和存储密钥
	global.GVA_LOG.Warn("SSH主机密钥不存在或无效，请使用 ssh-keygen 生成密钥")
	global.GVA_LOG.Info(fmt.Sprintf("提示: 运行以下命令生成密钥: ssh-keygen -f %s -t rsa", keyFile))

	return nil, fmt.Errorf("SSH主机密钥不存在")
}
