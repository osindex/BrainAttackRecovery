# BrainAttackRecovery 生产部署文档

本文档说明如何在服务器上部署本项目。部署假设：

- PostgreSQL 已经安装在**宿主机**上
- PostgreSQL 监听宿主机 `5432`
- Docker 只运行两个容器：`linapro` 和 `h5`
- 不在 `docker-compose.prod.yml` 里再启动 PostgreSQL 容器
- 两个容器加入已存在的 `1panel-network`
- 只对外暴露 H5 入口 `18080`

---

## 1. 推荐目录结构

服务器上建议使用：

```text
/opt/brain-rehab/
├── docker-compose.prod.yml
├── .env.prod
├── infra/
│   └── linapro.config.yaml
├── deploy/
│   └── nginx.prod.conf
├── data/
│   ├── linapro-upload/
│   └── linapro-output/
└── backup/
```

创建目录：

```bash
sudo mkdir -p /opt/brain-rehab
sudo chown -R $USER:$USER /opt/brain-rehab
cd /opt/brain-rehab
mkdir -p infra data/linapro-upload data/linapro-output backup
```

---

## 2. 准备宿主机 PostgreSQL

本项目生产部署使用宿主机 PostgreSQL，不使用 PG 容器。

确认 PostgreSQL 可用：

```bash
psql -h 127.0.0.1 -p 5432 -U postgres -c "SELECT version();"
```

创建数据库：

```bash
createdb -h 127.0.0.1 -p 5432 -U postgres linapro
```

如果数据库已经存在，可以跳过。

如果 PostgreSQL 只监听 `127.0.0.1`，Docker 容器可能无法通过 `host.docker.internal` 访问。请根据你的服务器环境检查：

```bash
sudo ss -ltnp | grep 5432
```

如果需要允许 Docker 网关访问，通常要调整 PostgreSQL：

```conf
# postgresql.conf
listen_addresses = '*'

# pg_hba.conf，示例网段按你的 Docker 网关实际网段调整
host    linapro    postgres    172.16.0.0/12    md5
```

修改后重启 PostgreSQL：

```bash
sudo systemctl restart postgresql
```

也可以用更严格的 Docker 网关 IP 替代 `172.16.0.0/12`。

> 注意：生产 compose 使用 1Panel 的 `1panel-network`，`linapro` 容器通过 `host.docker.internal:5432` 访问宿主机 PostgreSQL。请确保 PostgreSQL 允许 Docker 网关来源连接。

---

## 3. 准备部署文件

从仓库复制这些文件到 `/opt/brain-rehab/`：

```text
deploy/docker-compose.prod.yml  -> /opt/brain-rehab/docker-compose.prod.yml
deploy/.env.prod.example        -> /opt/brain-rehab/.env.prod
infra/linapro.config.yaml       -> /opt/brain-rehab/infra/linapro.config.yaml
deploy/nginx.prod.conf          -> /opt/brain-rehab/deploy/nginx.prod.conf
deploy/backup.sh                -> /opt/brain-rehab/backup.sh
deploy/restore.sh               -> /opt/brain-rehab/restore.sh
```

示例：

```bash
cp deploy/docker-compose.prod.yml /opt/brain-rehab/docker-compose.prod.yml
cp deploy/.env.prod.example /opt/brain-rehab/.env.prod
cp infra/linapro.config.yaml /opt/brain-rehab/infra/linapro.config.yaml
mkdir -p /opt/brain-rehab/deploy
cp deploy/nginx.prod.conf /opt/brain-rehab/deploy/nginx.prod.conf
cp deploy/backup.sh /opt/brain-rehab/backup.sh
cp deploy/restore.sh /opt/brain-rehab/restore.sh
chmod +x /opt/brain-rehab/backup.sh /opt/brain-rehab/restore.sh
```

---

## 4. 配置 `.env.prod`

编辑：

```bash
nano /opt/brain-rehab/.env.prod
```

内容示例：

```env
H5_IMAGE=ghcr.io/<owner>/<repo>/brain-rehab-h5:main
LINAPRO_IMAGE=ghcr.io/<owner>/<repo>/brain-rehab-linapro:main

POSTGRES_DSN=pgsql:postgres:<你的PostgreSQL密码>@tcp(host.docker.internal:5432)/linapro?sslmode=disable
LINAPRO_JWT_SECRET=<至少32位随机字符串>
```

生成 JWT secret：

```bash
openssl rand -base64 32
```

如果你的 GHCR 镜像是私有的，需要先登录：

```bash
echo <你的GitHub PAT> | docker login ghcr.io -u <你的GitHub用户名> --password-stdin
```

PAT 至少需要 `read:packages` 权限。

---

## 5. 启动服务

先确认 1Panel 网络已存在：

```bash
docker network inspect 1panel-network >/dev/null
```

如果不存在，请先在 1Panel 中创建，或者执行：

```bash
docker network create 1panel-network
```

```bash
cd /opt/brain-rehab
docker compose --env-file .env.prod -f docker-compose.prod.yml pull
docker compose --env-file .env.prod -f docker-compose.prod.yml up -d
```

查看状态：

```bash
docker compose --env-file .env.prod -f docker-compose.prod.yml ps
```

查看日志：

```bash
docker compose --env-file .env.prod -f docker-compose.prod.yml logs -f linapro
docker compose --env-file .env.prod -f docker-compose.prod.yml logs -f h5
```

---

## 6. 访问地址

生产 compose 会把 `linapro` 和 `h5` 加入已存在的 `1panel-network`。只有 `h5` 对外暴露 `18080`，`linapro` 不直接暴露宿主机端口。

注意：`h5` 容器内部 nginx 监听 `80`，宿主机通过 compose 映射为 `18080:80`。所以不要把 `deploy/nginx.prod.conf` 里的 `listen` 改成 `18080`。

生产 nginx 使用 `deploy/nginx.prod.conf`，把 `/api/*` 反代到同一 Docker 网络内的 `linapro:8080`。

本机访问：

```text
http://127.0.0.1:18080/
```

局域网访问：

```text
http://服务器IP:18080/
```

API 健康检查：

```bash
curl http://127.0.0.1:18080/api/v1/health
```

期望返回：

```json
{"code":0,"message":"OK","data":{"status":"ok","mode":"single"}}
```

---

## 7. 防火墙建议

如果只在服务器本机访问，不开放端口即可。

如果局域网手机需要访问，开放 `18080`：

```bash
sudo ufw allow 18080/tcp
```

不建议开放 PostgreSQL `5432` 到公网。

---

## 8. 备份

数据库在宿主机 PostgreSQL 内，上传文件在 `/opt/brain-rehab/data/`。

备份：

```bash
cd /opt/brain-rehab
./backup.sh
```

备份内容：

- `backup/linapro-时间.sql`
- `backup/linapro-files-时间.tar.gz`

建议加 cron 每天备份：

```bash
crontab -e
```

加入：

```cron
30 2 * * * cd /opt/brain-rehab && ./backup.sh >> backup/backup.log 2>&1
```

---

## 9. 恢复

```bash
cd /opt/brain-rehab
./restore.sh backup/linapro-2026-05-09-020000.sql backup/linapro-files-2026-05-09-020000.tar.gz
```

如果只恢复数据库：

```bash
./restore.sh backup/linapro-2026-05-09-020000.sql
```

---

## 10. 更新版本

当 GitHub Actions 推送了新镜像后，在服务器执行：

```bash
cd /opt/brain-rehab
docker compose --env-file .env.prod -f docker-compose.prod.yml pull
docker compose --env-file .env.prod -f docker-compose.prod.yml up -d
```

---

## 11. 常见问题

### 1. `linapro` 连不上数据库

检查：

```bash
psql -h 127.0.0.1 -p 5432 -U postgres linapro
```

再检查 `.env.prod`：

```env
POSTGRES_DSN=pgsql:postgres:<密码>@tcp(host.docker.internal:5432)/linapro?sslmode=disable
```

### 2. `18080` 端口被占用

查看占用：

```bash
sudo ss -ltnp | grep 18080
```

如果要改端口，需要修改 H5 nginx 镜像或改为前置 Caddy/Nginx 反代。当前推荐保持 `18080`。

### 3. GHCR 拉不到镜像

如果镜像是私有包，先登录：

```bash
echo <PAT> | docker login ghcr.io -u <GitHub用户名> --password-stdin
```

如果不想登录，把 GitHub Packages 中两个镜像改成 public。
