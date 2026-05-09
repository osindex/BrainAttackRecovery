-- Mock picture-card data for local validation.
-- 本地验证使用的图卡演示数据。

INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '苹果', '苹果', 'https://commons.wikimedia.org/wiki/Special:FilePath/Red_Apple.jpg', 'wikimedia', 'Wikimedia Commons', 1, 1, 10, 'Real object photo mock card'
FROM plugin_rehab_card_category c
WHERE c."code" = 'food'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '苹果' AND card."deleted_at" IS NULL);

INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '杯子', '杯子', 'https://commons.wikimedia.org/wiki/Special:FilePath/White_cup_and_saucer.jpg', 'wikimedia', 'Wikimedia Commons', 1, 1, 20, 'Real object photo mock card'
FROM plugin_rehab_card_category c
WHERE c."code" = 'daily-object'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '杯子' AND card."deleted_at" IS NULL);
