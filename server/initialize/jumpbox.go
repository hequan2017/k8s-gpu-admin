package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service/k8sgpu/jumpbox"
	"go.uber.org/zap"
)

const (
	// JumpboxPort SSH跳板机监听端口
	JumpboxPort = 2026
)

// InitJumpbox 初始化SSH跳板机
func InitJumpbox() {
	err := jumpbox.JumpboxServiceApp.Start(JumpboxPort)
	if err != nil {
		global.GVA_LOG.Error("SSH跳板机启动失败", zap.Error(err))
		// 不阻塞主程序启动
		return
	}
}
