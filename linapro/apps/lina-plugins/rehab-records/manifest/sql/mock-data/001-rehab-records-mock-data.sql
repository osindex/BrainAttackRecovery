-- Mock training records for local validation.
-- 本地验证使用的训练记录演示数据。

INSERT INTO plugin_rehab_record ("client_record_id", "training_type", "occurred_on", "duration_seconds", "payload", "remark") VALUES
('mock-walk-001', 'walk', CURRENT_DATE, 600, '{}', 'Mock walking record'),
('mock-fist-001', 'fist_raise', CURRENT_DATE, 0, '{"sets":2,"reps":10}', 'Mock fist raise record'),
('mock-gaze-001', 'eye_gaze', CURRENT_DATE, 0, '{"count":20}', 'Mock eye gaze record')
ON CONFLICT DO NOTHING;
