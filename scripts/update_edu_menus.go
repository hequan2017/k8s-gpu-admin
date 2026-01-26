package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MenuItem 菜单结构
type MenuItem struct {
	ID       uint   `gorm:"column:id;primaryKey"`
	ParentID *uint  `gorm:"column:parent_id"`
	Path     string `gorm:"column:path"`
	Name     string `gorm:"column:name"`
	Sort     int    `gorm:"column:sort"`
}

func main() {
	// 请根据实际配置修改数据库连接信息
	dsn := "root:123456@tcp(127.0.0.1:3306)/gva?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}

	fmt.Println("数据库连接成功!")

	// 更新所有教育平台子菜单的父节点为 104 (天启智教)
	result := db.Exec("UPDATE sys_base_menus SET parent_id = 104 WHERE id IN (94, 95, 96, 97, 98, 99, 100, 101, 102, 103)")
	if result.Error != nil {
		fmt.Println("更新父节点失败:", result.Error)
		return
	}
	fmt.Printf("成功更新 %d 条记录的父节点\n", result.RowsAffected)

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
		db.Exec("UPDATE sys_base_menus SET sort = ? WHERE id = ?", u.Sort, u.ID)
	}

	fmt.Println("菜单排序更新完成!")
	fmt.Println("\n✅ 教育平台菜单结构已组织到父菜单'天启智教'下")
}
