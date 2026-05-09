# 设计说明：卒中康复训练 MVP

## 架构

- `h5/`：患者移动端 H5，使用 Vue 3、Vant、Pinia、Dexie、PWA。
- `linapro/apps/lina-plugins/rehab-cards`：图卡管理源码插件。
- `linapro/apps/lina-plugins/rehab-records`：训练记录源码插件。
- `docker-compose.yml`：本地部署 PostgreSQL、LinaPro 和 H5 nginx。

## 数据同步

H5 训练记录先写入 IndexedDB，状态为 `pending`。同步时调用 `POST /api/v1/rehab/record`。后端以 `clientRecordId` 作为幂等键，重复提交返回已有记录 ID。

## 安全边界

H5 不内置管理员凭据。设备需要在 **本机配对** 页面输入低权限 LinaPro 账号。Compose 默认只绑定本机回环地址。JWT secret、PostgreSQL 密码和 DSN 通过 `.env` 注入。

## 已知折中

- 插件服务层当前仍使用 GoFrame `g.DB().Model` 的轻量实现，后续应补齐 `gf gen dao` 生成的 DAO/DO/Entity 并替换为项目标准数据访问模式。
- 管理后台页面已具备基本 CRUD/执行能力，但尚未完全采用 Vben Grid/Form 规范。
- 图片爬取为占位图生成，接外部 API 前需要补充 provider allowlist、下载限流、文件校验和 SSRF 防护。
