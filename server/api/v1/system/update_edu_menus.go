package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/gin-gonic/gin"
)

// UpdateEduMenus 更新教育平台菜单结构
// @Tags System
// @Summary 更新教育平台菜单结构
// @Description 将所有教育平台子菜单组织到父菜单"天启智教"下
// @Success 200 {object} response.Response{msg=string}
// @Router /sys/updateEduMenus [post]
func (s *SystemApi) UpdateEduMenus(c *gin.Context) {
	// 更新所有教育平台子菜单的父节点为 104 (天启智教)
	result := global.GVA_DB.Exec("UPDATE sys_base_menus SET parent_id = 104 WHERE id IN (94, 95, 96, 97, 98, 99, 100, 101, 102, 103)")
	if result.Error != nil {
		common.FailWithMessage(c, "更新父节点失败: "+result.Error.Error())
		return
	}

	// 更新排序
	updates := []struct {
		ID   uint
		Sort int
	}{
		{94, 1},  // 课程
		{95, 2},  // 课程章节
		{96, 3},  // 课程课时
		{97, 4},  // 实训项目
		{98, 5},  // 学员作品
		{99, 6},  // AI创作
		{100, 7}, // 竞赛活动
		{101, 8}, // 考证管理
		{102, 9}, // 技能树
		{103, 10}, // 就业服务
	}

	for _, u := range updates {
		global.GVA_DB.Exec("UPDATE sys_base_menus SET sort = ? WHERE id = ?", u.Sort, u.ID)
	}

	common.OkWithMessage(c, "教育平台菜单结构已更新")
}
