-- Mock picture-card data for local validation.
-- 本地验证使用的图卡演示数据。

INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '苹果', '苹果', 'https://dummyimage.com/512x512/f5f5f5/333333&text=%E8%8B%B9%E6%9E%9C', 'mock', 'placeholder', 1, 1, 10, 'Mock card'
FROM plugin_rehab_card_category c
WHERE c."code" = 'food'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '苹果' AND card."deleted_at" IS NULL);

INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '杯子', '杯子', 'https://dummyimage.com/512x512/f5f5f5/333333&text=%E6%9D%AF%E5%AD%90', 'mock', 'placeholder', 1, 1, 20, 'Mock card'
FROM plugin_rehab_card_category c
WHERE c."code" = 'daily-object'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '杯子' AND card."deleted_at" IS NULL);
