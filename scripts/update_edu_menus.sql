-- ============================================
-- 天启智教平台菜单结构调整SQL脚本
-- ============================================
-- 说明：将所有教育平台子菜单组织到父菜单"天启智教"(ID: 104)下
-- 执行方式：在数据库管理工具中直接执行
-- ============================================

-- 1. 更新所有教育平台子菜单的父节点为 104 (天启智教)
UPDATE sys_base_menus SET parent_id = 104 WHERE id IN (94, 95, 96, 97, 98, 99, 100, 101, 102, 103);

-- 2. 更新菜单排序顺序
UPDATE sys_base_menus SET sort = 1 WHERE id = 94;   -- 课程
UPDATE sys_base_menus SET sort = 2 WHERE id = 95;   -- 课程章节
UPDATE sys_base_menus SET sort = 3 WHERE id = 96;   -- 课程课时
UPDATE sys_base_menus SET sort = 4 WHERE id = 97;   -- 实训项目
UPDATE sys_base_menus SET sort = 5 WHERE id = 98;   -- 学员作品
UPDATE sys_base_menus SET sort = 6 WHERE id = 99;   -- AI创作
UPDATE sys_base_menus SET sort = 7 WHERE id = 100;  -- 竞赛活动
UPDATE sys_base_menus SET sort = 8 WHERE id = 101;  -- 考证管理
UPDATE sys_base_menus SET sort = 9 WHERE id = 102;  -- 技能树
UPDATE sys_base_menus SET sort = 10 WHERE id = 103;  -- 就业服务

-- 3. 验证更新结果
SELECT id, parent_id, title, path, sort
FROM sys_base_menus
WHERE id = 104 OR parent_id = 104
ORDER BY parent_id, sort;

-- ============================================
-- 执行完成后，菜单结构如下：
-- 📚 天启智教 (104)
-- ├── 📖 课程 (94)
-- ├── 📑 课程章节 (95)
-- ├── 🎬 课程课时 (96)
-- ├── 💻 实训项目 (97)
-- ├── 🎨 学员作品 (98)
-- ├── 🤖 AI创作 (99)
-- ├── 🏆 竞赛活动 (100)
-- ├── 📜 考证管理 (101)
-- ├── 🌳 技能树 (102)
-- └── 💼 就业服务 (103)
-- ============================================
