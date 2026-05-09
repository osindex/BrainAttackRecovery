-- 001: rehab-cards schema
-- 001：康复图卡管理数据结构

CREATE TABLE IF NOT EXISTS plugin_rehab_card_category (
    "id"          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "name"        VARCHAR(100) NOT NULL DEFAULT '',
    "code"        VARCHAR(100) NOT NULL DEFAULT '',
    "sort"        INTEGER      NOT NULL DEFAULT 0,
    "status"      SMALLINT     NOT NULL DEFAULT 1,
    "remark"      VARCHAR(500) NOT NULL DEFAULT '',
    "created_by"  BIGINT       NOT NULL DEFAULT 0,
    "updated_by"  BIGINT       NOT NULL DEFAULT 0,
    "created_at"  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at"  TIMESTAMP    NULL DEFAULT NULL
);

COMMENT ON TABLE plugin_rehab_card_category IS 'Stroke rehabilitation picture-card category table';
COMMENT ON COLUMN plugin_rehab_card_category."name" IS 'Category display name';
COMMENT ON COLUMN plugin_rehab_card_category."code" IS 'Category stable code';
COMMENT ON COLUMN plugin_rehab_card_category."status" IS 'Status: 1=enabled, 0=disabled';

CREATE UNIQUE INDEX IF NOT EXISTS uk_plugin_rehab_card_category_code ON plugin_rehab_card_category ("code");
CREATE INDEX IF NOT EXISTS idx_plugin_rehab_card_category_status ON plugin_rehab_card_category ("status");

CREATE TABLE IF NOT EXISTS plugin_rehab_card (
    "id"             BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "category_id"    BIGINT       NOT NULL DEFAULT 0,
    "title"          VARCHAR(100) NOT NULL DEFAULT '',
    "label"          VARCHAR(100) NOT NULL DEFAULT '',
    "image_url"      VARCHAR(500) NOT NULL DEFAULT '',
    "image_file_id"  BIGINT       NOT NULL DEFAULT 0,
    "source"         VARCHAR(100) NOT NULL DEFAULT 'manual',
    "source_url"     VARCHAR(500) NOT NULL DEFAULT '',
    "license"        VARCHAR(100) NOT NULL DEFAULT '',
    "difficulty"     SMALLINT     NOT NULL DEFAULT 1,
    "status"         SMALLINT     NOT NULL DEFAULT 1,
    "sort"           INTEGER      NOT NULL DEFAULT 0,
    "remark"         VARCHAR(500) NOT NULL DEFAULT '',
    "created_by"     BIGINT       NOT NULL DEFAULT 0,
    "updated_by"     BIGINT       NOT NULL DEFAULT 0,
    "created_at"     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at"     TIMESTAMP    NULL DEFAULT NULL
);

COMMENT ON TABLE plugin_rehab_card IS 'Stroke rehabilitation picture-card table';
COMMENT ON COLUMN plugin_rehab_card."category_id" IS 'Picture-card category ID';
COMMENT ON COLUMN plugin_rehab_card."title" IS 'Card title';
COMMENT ON COLUMN plugin_rehab_card."label" IS 'Expected answer label';
COMMENT ON COLUMN plugin_rehab_card."difficulty" IS 'Difficulty: 1=easy, 2=normal, 3=hard';
COMMENT ON COLUMN plugin_rehab_card."status" IS 'Status: 1=enabled, 0=disabled';

CREATE INDEX IF NOT EXISTS idx_plugin_rehab_card_category ON plugin_rehab_card ("category_id");
CREATE INDEX IF NOT EXISTS idx_plugin_rehab_card_status ON plugin_rehab_card ("status");

CREATE TABLE IF NOT EXISTS plugin_rehab_card_crawl_job (
    "id"             BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "category_id"    BIGINT       NOT NULL DEFAULT 0,
    "keyword"        VARCHAR(100) NOT NULL DEFAULT '',
    "provider"       VARCHAR(100) NOT NULL DEFAULT 'manual',
    "requested_count" INTEGER     NOT NULL DEFAULT 0,
    "fetched_count"   INTEGER     NOT NULL DEFAULT 0,
    "status"         VARCHAR(32)  NOT NULL DEFAULT 'pending',
    "message"        VARCHAR(500) NOT NULL DEFAULT '',
    "created_by"     BIGINT       NOT NULL DEFAULT 0,
    "updated_by"     BIGINT       NOT NULL DEFAULT 0,
    "created_at"     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at"     TIMESTAMP    NULL DEFAULT NULL
);

COMMENT ON TABLE plugin_rehab_card_crawl_job IS 'Picture-card crawler job table';
COMMENT ON COLUMN plugin_rehab_card_crawl_job."status" IS 'Job status: pending, completed, failed';

CREATE INDEX IF NOT EXISTS idx_plugin_rehab_card_crawl_job_category ON plugin_rehab_card_crawl_job ("category_id");
CREATE INDEX IF NOT EXISTS idx_plugin_rehab_card_crawl_job_status ON plugin_rehab_card_crawl_job ("status");

INSERT INTO plugin_rehab_card_category ("name", "code", "sort", "status", "remark") VALUES
('动物', 'animal', 10, 1, '常见动物命名训练'),
('食物', 'food', 20, 1, '常见食物命名训练'),
('日用品', 'daily-object', 30, 1, '日常用品命名训练'),
('交通工具', 'vehicle', 40, 1, '交通工具命名训练'),
('衣物', 'clothing', 50, 1, '衣物命名训练'),
('身体部位', 'body-part', 60, 1, '身体部位命名训练')
ON CONFLICT DO NOTHING;
