# 本地基础设施

该目录保存根目录 `docker-compose.yml` 使用的本地运行配置。

## 服务

- PostgreSQL 16：`127.0.0.1:15432`
- 统一本地入口：`http://127.0.0.1:18080`
- 患者 H5：`/`
- LinaPro API：`/api/*` 反代

LinaPro 容器通过 `GF_GCFG_PATH=/app` 读取 `infra/linapro.config.yaml`。

运行前请复制根目录 `.env.example` 为 `.env` 并替换所有占位密钥。
