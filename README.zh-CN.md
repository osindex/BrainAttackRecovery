# BrainAttackRecovery

面向单个患者自用的卒中康复训练自托管应用。

## 目录结构

```text
.
├── docker-compose.yml        # 本地服务栈：PostgreSQL + LinaPro + H5 nginx
├── infra/                    # 本地 LinaPro 运行配置
├── h5/                       # 患者使用的 H5 应用（Vue 3 + Vant + PWA）
└── linapro/                  # vendored LinaPro 源码 fork 与康复插件
```

## 当前 MVP

- 患者 H5 包含慢走、握拳平举、眼睛凝视、图片卡辨物和历史记录页面。
- 训练记录优先写入 IndexedDB，本机配对后通过同步队列上传到 LinaPro。
- LinaPro 源码插件：
  - `rehab-cards`：图卡分类、图卡、占位爬取任务。
  - `rehab-records`：训练记录与按日聚合。
- 本地 Docker 服务栈使用 PostgreSQL，默认只绑定到 `127.0.0.1`。
- 统一本地入口：`http://127.0.0.1:18080`。
- 患者 H5 位于 `/`；LinaPro API 通过 `/api/*` 反代。

## 本地密钥

运行 Docker Compose 前，复制 `.env.example` 为 `.env` 并替换全部占位值。

```powershell
copy .env.example .env
docker compose up --build
```

## 安全说明

- H5 不再内置默认管理员凭据。
- 请在 H5 的 **本机配对** 页面输入为该设备创建的低权限 LinaPro 账号。
- 不要把内置 `admin/admin123` 账号用于患者设备。
- `docker-compose.yml` 默认只把服务端口绑定到本机回环地址，便于本地开发。

## 验证

本地已验证：

- 在 `linapro/` 下执行 `go build ./apps/lina-core`
- 在 `h5/` 下执行 `pnpm build`
- 在根目录执行 `docker compose config`

未登录 Docker Hub 的环境可能因为基础镜像拉取触发 HTTP 429 限流，导致 `docker compose build` 失败。

## GitHub 容器打包

GitHub Actions 工作流 `.github/workflows/docker-images.yml` 会构建两个镜像：

- `ghcr.io/<owner>/<repo>/brain-rehab-h5`
- `ghcr.io/<owner>/<repo>/brain-rehab-linapro`

Pull Request 只构建不推送；推送到 `main` 或版本 tag 时会推送到 GHCR。

本地 H5 Dockerfile 使用已经构建好的 `h5/dist`，用于绕开本机 Node 镜像版本不一致问题。CI 使用 `h5/Dockerfile.ci`，会在容器内执行干净的 Node 构建。
