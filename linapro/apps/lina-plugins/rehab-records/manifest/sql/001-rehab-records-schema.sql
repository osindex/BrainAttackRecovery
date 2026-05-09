-- 001: rehab-records schema
-- 001：康复训练记录数据结构

CREATE TABLE IF NOT EXISTS plugin_rehab_record (
    "id"                BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "client_record_id"  VARCHAR(100) NOT NULL DEFAULT '',
    "training_type"     VARCHAR(50)  NOT NULL DEFAULT '',
    "occurred_on"       DATE         NOT NULL DEFAULT CURRENT_DATE,
    "started_at"        TIMESTAMP    NULL DEFAULT NULL,
    "ended_at"          TIMESTAMP    NULL DEFAULT NULL,
    "duration_seconds"  INTEGER      NOT NULL DEFAULT 0,
    "sets_count"        INTEGER      NOT NULL DEFAULT 0,
    "reps_count"        INTEGER      NOT NULL DEFAULT 0,
    "gaze_count"        INTEGER      NOT NULL DEFAULT 0,
    "card_id"           BIGINT       NOT NULL DEFAULT 0,
    "is_correct"        SMALLINT     NOT NULL DEFAULT 0,
    "reaction_ms"       INTEGER      NOT NULL DEFAULT 0,
    "payload"           TEXT         NOT NULL DEFAULT '{}',
    "remark"            VARCHAR(500) NOT NULL DEFAULT '',
    "created_by"        BIGINT       NOT NULL DEFAULT 0,
    "updated_by"        BIGINT       NOT NULL DEFAULT 0,
    "created_at"        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at"        TIMESTAMP    NULL DEFAULT NULL
);

COMMENT ON TABLE plugin_rehab_record IS 'Stroke rehabilitation training record table';
COMMENT ON COLUMN plugin_rehab_record."client_record_id" IS 'Client-generated idempotency key';
COMMENT ON COLUMN plugin_rehab_record."training_type" IS 'Training type: walk, fist_raise, eye_gaze, card_game';
COMMENT ON COLUMN plugin_rehab_record."occurred_on" IS 'Training date';
COMMENT ON COLUMN plugin_rehab_record."duration_seconds" IS 'Duration in seconds for timer-based exercises';
COMMENT ON COLUMN plugin_rehab_record."sets_count" IS 'Set count for fist-raise exercise';
COMMENT ON COLUMN plugin_rehab_record."reps_count" IS 'Repetition count for fist-raise exercise';
COMMENT ON COLUMN plugin_rehab_record."gaze_count" IS 'Left-right gaze repetition count';
COMMENT ON COLUMN plugin_rehab_record."card_id" IS 'Picture-card id for card game records';
COMMENT ON COLUMN plugin_rehab_record."is_correct" IS 'Card game correctness: 1=correct, 0=incorrect';
COMMENT ON COLUMN plugin_rehab_record."reaction_ms" IS 'Card game reaction time in milliseconds';
COMMENT ON COLUMN plugin_rehab_record."payload" IS 'Additional JSON payload stored as text for MVP portability';

CREATE UNIQUE INDEX IF NOT EXISTS uk_plugin_rehab_record_client_id ON plugin_rehab_record ("client_record_id");
CREATE INDEX IF NOT EXISTS idx_plugin_rehab_record_type_date ON plugin_rehab_record ("training_type", "occurred_on");
CREATE INDEX IF NOT EXISTS idx_plugin_rehab_record_created_by ON plugin_rehab_record ("created_by");
