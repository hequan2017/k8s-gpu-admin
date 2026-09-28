[简体中文](README.md) | [English](README.en.md)

# k8s-gpu-admin

A GPU compute management platform built on [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin): centrally manage GPU compute nodes, provision Docker container instances with GPU passthrough in one click by spec, with a built-in SSH jumpbox and MCP toolset.

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go&logoColor=white)
![Vue](https://img.shields.io/badge/Vue-3.x-4FC08D?logo=vuedotjs&logoColor=white)
![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)

## Introduction

When you own multiple GPU servers — especially ones offering compute through Docker — the usual pain points are: machine details scattered across spreadsheets, manual SSH commands for every delivery, and resource allocation kept in someone's head, which often leads to the same GPU card being handed out twice (overselling).

k8s-gpu-admin moves this whole workflow to the web: administrators register GPU servers (CPU / memory / disks / GPU model and count, SSH credentials, remote Docker endpoint) and publish them, define sellable product specs (e.g. "2 x A100 / 16C / 64G"), and maintain a shared image registry. When a user creates an instance from an image and a spec, the platform automatically picks a node with enough remaining resources, pulls the image via the remote Docker API, passes through the requested number of GPUs with NVIDIA DeviceRequests, applies CPU / memory limits, and mounts a data disk — a one-stop delivery.

The platform also inherits the full gin-vue-admin v2.8.8 foundation (users / roles / Casbin authorization, menus, API management, code generator, form designer, plugin system, etc.), plus a built-in SSH jumpbox and an MCP Server for AI coding assistants. It fits individuals or small teams managing self-built GPU compute pools and renting compute internally.

## ✨ Features

- **Compute node management**: register CPU / memory / system disk / data disk / GPU model and count, public & private IPs, SSH port and credentials, remote Docker endpoint (TLS certificate supported); publish / unpublish nodes
- **Product spec management**: define GPU model / count, CPU, memory, disks and hourly pricing; published specs become available for provisioning
- **Image registry management**: maintain image names, addresses and sources; published images can be selected when creating instances
- **One-click instance provisioning**: automatically matches a published node with enough remaining resources and creates the container remotely; GPUs are passed through via NVIDIA DeviceRequests, CPU / memory are limited, the system disk is capped with overlay2 `storage-opt` (automatically skipped when the storage driver does not support it), and the data disk is mounted as a dedicated named volume at `/data`
- **Instance lifecycle management**: start / stop / restart containers, view container logs, sync container status; deleting an instance also cleans up its container and data volume
- **Resource accounting to prevent overselling**: tracks per-node CPU / memory / GPU / disk usage aggregated from instance specs, and only nodes with sufficient headroom are eligible for allocation
- **Per-user data isolation**: regular users can only view and operate their own instances; admins (authority ID 888) manage everything
- **SSH jumpbox**: a built-in SSH Server (default port 2026) that authenticates directly against platform accounts (bcrypt check), lists the instances you are allowed to operate, and lets you pick one to connect to (container terminal integration in progress)
- **MCP Server**: bundled MCP tools (API / menu generation, requirement analysis, code review, dictionary queries, etc.), exposed via SSE (default `/sse`, `/message`) for AI coding assistants
- **Full admin foundation**: inherits gin-vue-admin's users / roles / Casbin authorization, menu & API management, code generator, form designer, plugins (email, announcement), multi-cloud object storage uploads, operation logs, and Swagger docs

## 🛠 Tech Stack

| Layer | Technologies |
| --- | --- |
| Backend | Go 1.24, Gin 1.10, GORM 1.25 (MySQL / PostgreSQL / SQLite / SQLServer / Oracle), Casbin v2, JWT v5, Zap, Viper, Go-Redis v9, Qmgo (MongoDB), Docker SDK, mcp-go, Swaggo |
| Frontend | Vue 3.5, Vite, Element Plus 2.10, Pinia 2, Axios, ECharts, UnoCSS |
| Deployment | Docker, docker compose, Kubernetes manifests, Makefile, nginx |

## 🚀 Quick Start

### Option 1: One-command start with docker compose

```bash
cd deploy/docker-compose
docker compose up -d
```

This starts 4 containers: frontend web (port `8080`), backend server (port `8888`), MySQL 8 (host port `13306`) and Redis 6 (host port `16379`), with data persisted in the `mysql` and `redis` volumes.

Open `http://localhost:8080` and log in with the default account `admin / 123456` (seeded by `server/source/system/user.go`; change it after the first login).

### Option 2: Local development

```bash
# Backend: listens on port 8888 by default, config file is server/config.yaml
cd server
go mod tidy
go run .

# Frontend (in another terminal)
cd web
yarn install
yarn dev
```

### Key configuration (server/config.yaml)

| Item | Description |
| --- | --- |
| `system.db-type` / `system.addr` | Database type (default mysql) and listen address (default `:8888`) |
| `mysql` / `redis` / `mongo` | Database connections; `system.use-redis` / `use-mongo` toggle them |
| `jwt.signing-key` | JWT signing key — change it in production |
| `mcp.addr` / `mcp.sse_path` / `mcp.message_path` | MCP Server port (default 8889) and SSE paths |
| SSH jumpbox port | Constant defined in `server/initialize/jumpbox.go` (default 2026) |

### Production build & deployment

```bash
make build            # Build frontend & backend in containers, output to build/
make image            # Build the combined frontend+backend image
make doc              # Regenerate Swagger docs (requires swag)

# Kubernetes deployment (ConfigMap / Deployment / Service / Ingress included)
kubectl apply -f deploy/kubernetes/server/
kubectl apply -f deploy/kubernetes/web/
```

## 📁 Directory Structure

```
├── server/                    # Go backend (gin-vue-admin server + k8sgpu business module)
│   ├── api/v1/k8sgpu/         # Node / spec / image / instance APIs
│   ├── service/k8sgpu/        # Instance orchestration, resource accounting, SSH jumpbox
│   ├── mcp/                   # MCP toolset
│   ├── model/k8sgpu/          # Data models
│   └── config.yaml            # Backend configuration
├── web/                       # Vue 3 frontend
│   └── src/view/k8sgpu/       # Node / spec / image / instance management pages
├── deploy/
│   ├── docker-compose/        # One-command compose stack
│   └── kubernetes/            # Kubernetes manifests
├── scripts/                   # Helper scripts
└── Makefile                   # Build / package / docs commands
```

## 🔗 Related Projects

- [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin): the admin foundation this project is built on (v2.8.8)

## 📄 License

Licensed under the [Apache License 2.0](LICENSE).
