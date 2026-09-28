[简体中文](README.md) | [English](README.en.md)

# k8s-gpu-admin

基于 [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin) 的 GPU 算力管理平台：统一纳管 GPU 算力节点，按规格一键开通带 GPU 直通的 Docker 容器实例，内置 SSH 跳板机与 MCP 工具集。

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go&logoColor=white)
![Vue](https://img.shields.io/badge/Vue-3.x-4FC08D?logo=vuedotjs&logoColor=white)
![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)

## 项目介绍

在拥有多台 GPU 服务器（尤其是以 Docker 对外提供算力的场景）时，常见痛点是：机器信息散落在表格里、开机交付靠人工 SSH 敲命令、资源分配全凭记忆，容易出现同一张卡被重复分配的"超卖"问题。

k8s-gpu-admin 把这套流程搬到 Web 上：管理员把 GPU 服务器的 CPU / 内存 / 磁盘 / 显卡型号与数量、SSH 信息、远程 Docker 连接地址录入平台并上架，再定义可售卖的产品规格（如"2 x A100 / 16C / 64G"）和维护公共镜像库。普通用户选择镜像和规格创建实例后，平台会自动挑一个剩余资源足够的节点，调用远端 Docker API 拉镜像、以 NVIDIA DeviceRequests 直通指定数量的 GPU、限制 CPU / 内存并挂载数据盘，一条龙完成交付。

平台同时继承 gin-vue-admin v2.8.8 的全部基座能力（用户 / 角色 / Casbin 权限、菜单、API 管理、代码生成器、表单设计器、插件机制等），并内置了 SSH 跳板机和面向 AI 编码助手的 MCP Server。适合个人或小团队管理自建 GPU 算力池、做内部算力租用与分发。

## ✨ 功能特性

- **算力节点管理**：录入节点的 CPU / 内存 / 系统盘 / 数据盘 / 显卡型号与数量、公网内网 IP、SSH 端口与账号、远程 Docker 连接地址（支持 TLS 证书配置），支持上架 / 下架
- **产品规格管理**：定义显卡型号 / 数量、CPU、内存、磁盘与按小时计价，上架后可供用户选择开通
- **镜像库管理**：维护容器镜像名称、地址与来源，上架后供实例创建时选用
- **一键开通实例**：按规格自动匹配剩余资源足够的上架节点，远端拉取镜像并创建容器；GPU 以 NVIDIA DeviceRequests 直通，限制 CPU / 内存，系统盘通过 overlay2 `storage-opt` 限定（存储驱动不支持时自动降级忽略），数据盘挂独立命名卷到 `/data`
- **实例生命周期管理**：启动 / 停止 / 重启容器、查看容器日志、同步容器状态；删除实例时联动清理容器与数据卷
- **资源核算防超卖**：按实例规格累计各节点已用的 CPU / 内存 / GPU / 磁盘，只把有足够余量的节点纳入分配
- **数据权限隔离**：普通用户仅能查看和操作自己的实例，管理员（权限 ID 888）可管理全部数据
- **SSH 跳板机**：内置 SSH Server（默认端口 2026），直接用平台账号密码（bcrypt 校验）登录，按权限列出可操作实例并选择连接（容器终端对接开发中）
- **MCP Server**：内置 MCP 工具集（API / 菜单生成、需求分析、代码审查、字典查询等），以 SSE 方式（默认 `/sse`、`/message`）接入，供 AI 编码助手调用
- **完整后台基座**：继承 gin-vue-admin 的用户 / 角色 / Casbin 鉴权、菜单与 API 管理、代码生成器、表单设计器、插件（邮件、公告）、多云对象存储上传、操作日志、Swagger 文档

## 🛠 技术栈

| 端 | 技术 |
| --- | --- |
| 后端 | Go 1.24、Gin 1.10、GORM 1.25（MySQL / PostgreSQL / SQLite / SQLServer / Oracle）、Casbin v2、JWT v5、Zap、Viper、Go-Redis v9、Qmgo（MongoDB）、Docker SDK、mcp-go、Swaggo |
| 前端 | Vue 3.5、Vite、Element Plus 2.10、Pinia 2、Axios、ECharts、UnoCSS |
| 部署 | Docker、docker compose、Kubernetes 清单、Makefile、nginx |

## 🚀 快速开始

### 方式一：docker compose 一键启动

```bash
cd deploy/docker-compose
docker compose up -d
```

启动后包含 4 个容器：前端 web（`8080` 端口）、后端 server（`8888` 端口）、MySQL 8（宿主机 `13306`）、Redis 6（宿主机 `16379`），数据分别落在 `mysql`、`redis` 卷中持久化。

浏览器访问 `http://localhost:8080`，默认账号 `admin / 123456`（来源：`server/source/system/user.go` 初始化数据，首次登录后请及时修改）。

### 方式二：本地开发

```bash
# 后端：默认监听 8888 端口，配置文件为 server/config.yaml
cd server
go mod tidy
go run .

# 前端（另开终端）
cd web
yarn install
yarn dev
```

### 主要配置（server/config.yaml）

| 配置项 | 说明 |
| --- | --- |
| `system.db-type` / `system.addr` | 数据库类型（默认 mysql）、服务监听地址（默认 `:8888`） |
| `mysql` / `redis` / `mongo` | 数据库连接信息，`system.use-redis` / `use-mongo` 控制是否启用 |
| `jwt.signing-key` | JWT 签名密钥，生产环境务必修改 |
| `mcp.addr` / `mcp.sse_path` / `mcp.message_path` | MCP Server 端口（默认 8889）与 SSE 路径 |
| SSH 跳板机端口 | 常量定义在 `server/initialize/jumpbox.go`（默认 2026） |

### 生产构建与部署

```bash
make build            # 容器内构建前后端，产物输出到 build/
make image            # 打包前后端二合一镜像
make doc              # 重新生成 Swagger 文档（需安装 swag）

# Kubernetes 部署（含 ConfigMap / Deployment / Service / Ingress）
kubectl apply -f deploy/kubernetes/server/
kubectl apply -f deploy/kubernetes/web/
```

## 📁 目录结构

```
├── server/                    # Go 后端（gin-vue-admin server + k8sgpu 业务模块）
│   ├── api/v1/k8sgpu/         # 节点 / 规格 / 镜像 / 实例 API
│   ├── service/k8sgpu/        # 实例编排、资源核算、SSH 跳板机（jumpbox）
│   ├── mcp/                   # MCP 工具集
│   ├── model/k8sgpu/          # 数据模型
│   └── config.yaml            # 后端配置
├── web/                       # Vue 3 前端
│   └── src/view/k8sgpu/       # 节点 / 规格 / 镜像 / 实例管理页面
├── deploy/
│   ├── docker-compose/        # 一键 compose 编排
│   └── kubernetes/            # K8s 部署清单
├── scripts/                   # 辅助脚本
└── Makefile                   # 构建 / 打包 / 文档命令
```

## 🔗 相关项目

- [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)：本项目使用的后台管理基座（v2.8.8）

## 📄 License

本项目基于 [Apache License 2.0](LICENSE) 开源。
