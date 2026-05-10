-- Mock picture-card data for local validation.
-- 本地验证使用的图卡演示数据。
-- 图片统一通过 /api/v1/rehab/card/image/wiki 代理 Wikimedia Commons 实物图。

-- 食物类 food
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '苹果', '苹果', '/api/v1/rehab/card/image/wiki?file=Red_Apple.jpg', 'wikimedia', 'Wikimedia Commons', 1, 1, 10, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'food'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '苹果' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '香蕉', '香蕉', '/api/v1/rehab/card/image/wiki?file=Bananavarieties.jpg', 'wikimedia', 'Wikimedia Commons', 1, 1, 11, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'food'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '香蕉' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '橙子', '橙子', '/api/v1/rehab/card/image/wiki?file=Orange_fruit.jpg', 'wikimedia', 'Wikimedia Commons', 1, 1, 12, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'food'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '橙子' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '西瓜', '西瓜', '/api/v1/rehab/card/image/wiki?file=Watermelon_cross_BNC.jpg', 'wikimedia', 'Wikimedia Commons', 1, 1, 13, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'food'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '西瓜' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '草莓', '草莓', '/api/v1/rehab/card/image/wiki?file=Strawberry_Single1.jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 14, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'food'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '草莓' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '梨', '梨', '/api/v1/rehab/card/image/wiki?file=Yellow_pear_on_a_black_background.jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 15, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'food'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '梨' AND card."deleted_at" IS NULL);

-- 日用品类 daily-object
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '杯子', '杯子', '/api/v1/rehab/card/image/wiki?file=White_cup_and_saucer.jpg', 'wikimedia', 'Wikimedia Commons', 1, 1, 20, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'daily-object'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '杯子' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '勺子', '勺子', '/api/v1/rehab/card/image/wiki?file=Spoon_silver.jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 21, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'daily-object'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '勺子' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '筷子', '筷子', '/api/v1/rehab/card/image/wiki?file=Chopsticks-candc.jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 22, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'daily-object'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '筷子' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '水壶', '水壶', '/api/v1/rehab/card/image/wiki?file=WWII_Allied_Canteen.jpg', 'wikimedia', 'Wikimedia Commons', 3, 1, 23, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'daily-object'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '水壶' AND card."deleted_at" IS NULL);

-- 动物类 animal
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '小狗', '小狗', '/api/v1/rehab/card/image/wiki?file=Golden_Retriever_Carlos_(10581910556).jpg', 'wikimedia', 'Wikimedia Commons', 1, 1, 30, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'animal'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '小狗' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '小猫', '小猫', '/api/v1/rehab/card/image/wiki?file=Felis_catus-cat_on_snow.jpg', 'wikimedia', 'Wikimedia Commons', 1, 1, 31, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'animal'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '小猫' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '兔子', '兔子', '/api/v1/rehab/card/image/wiki?file=Oryctolagus_cuniculus_Rcdo.jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 32, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'animal'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '兔子' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '马', '马', '/api/v1/rehab/card/image/wiki?file=Hauspferd.JPG', 'wikimedia', 'Wikimedia Commons', 2, 1, 33, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'animal'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '马' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '鸡', '鸡', '/api/v1/rehab/card/image/wiki?file=Gallus_gallus_domesticus_Brown_Leghorn.jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 34, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'animal'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '鸡' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '牛', '牛', '/api/v1/rehab/card/image/wiki?file=Cow_female_black_white.jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 35, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'animal'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '牛' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '羊', '羊', '/api/v1/rehab/card/image/wiki?file=Sheep_in_field.JPG', 'wikimedia', 'Wikimedia Commons', 2, 1, 36, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'animal'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '羊' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '金鱼', '金鱼', '/api/v1/rehab/card/image/wiki?file=Goldfish3.jpg', 'wikimedia', 'Wikimedia Commons', 3, 1, 37, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'animal'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '金鱼' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '蝴蝶', '蝴蝶', '/api/v1/rehab/card/image/wiki?file=Monarch_In_May.jpg', 'wikimedia', 'Wikimedia Commons', 3, 1, 38, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'animal'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '蝴蝶' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '蚱蜢', '蚱蜢', '/api/v1/rehab/card/image/wiki?file=Grasshopper_1.jpg', 'wikimedia', 'Wikimedia Commons', 3, 1, 39, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'animal'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '蚱蜢' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '熊猫', '熊猫', '/api/v1/rehab/card/image/wiki?file=Panda_Cub_from_Wolong,_Sichuan,_China.JPG', 'wikimedia', 'Wikimedia Commons', 2, 1, 40, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'animal'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '熊猫' AND card."deleted_at" IS NULL);

-- 交通工具类 vehicle
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '汽车', '汽车', '/api/v1/rehab/card/image/wiki?file=2019_Toyota_Corolla_Icon_Tech_VVT-i_Hybrid_1.8.jpg', 'wikimedia', 'Wikimedia Commons', 1, 1, 50, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'vehicle'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '汽车' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '自行车', '自行车', '/api/v1/rehab/card/image/wiki?file=Bicycle_Abbey_Sprotbrough.jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 51, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'vehicle'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '自行车' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '公交车', '公交车', '/api/v1/rehab/card/image/wiki?file=Bus_in_Hong_Kong.jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 52, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'vehicle'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '公交车' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '飞机', '飞机', '/api/v1/rehab/card/image/wiki?file=Airbus_A380_blue_sky.jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 53, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'vehicle'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '飞机' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '救护车', '救护车', '/api/v1/rehab/card/image/wiki?file=Ambulance_Berlin.jpg', 'wikimedia', 'Wikimedia Commons', 3, 1, 54, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'vehicle'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '救护车' AND card."deleted_at" IS NULL);

-- 衣物类 clothing
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '鞋子', '鞋子', '/api/v1/rehab/card/image/wiki?file=Shoes_-_Nike_Air_Jordan_1_Retro_Banned_2016_-_sneakers.jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 60, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'clothing'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '鞋子' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '裤子', '裤子', '/api/v1/rehab/card/image/wiki?file=Blue_Denim_Jeans.jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 61, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'clothing'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '裤子' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '围巾', '围巾', '/api/v1/rehab/card/image/wiki?file=Red_scarf.jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 62, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'clothing'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '围巾' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '帽子', '帽子', '/api/v1/rehab/card/image/wiki?file=Winter_hat.jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 63, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'clothing'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '帽子' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '手套', '手套', '/api/v1/rehab/card/image/wiki?file=Glove.jpg', 'wikimedia', 'Wikimedia Commons', 3, 1, 64, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'clothing'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '手套' AND card."deleted_at" IS NULL);

-- 身体部位类 body-part
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '手', '手', '/api/v1/rehab/card/image/wiki?file=Hand_(sculpture).jpg', 'wikimedia', 'Wikimedia Commons', 1, 1, 70, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'body-part'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '手' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '眼睛', '眼睛', '/api/v1/rehab/card/image/wiki?file=Human_eye.jpg', 'wikimedia', 'Wikimedia Commons', 1, 1, 71, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'body-part'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '眼睛' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '脚', '脚', '/api/v1/rehab/card/image/wiki?file=Pair_of_feet.jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 72, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'body-part'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '脚' AND card."deleted_at" IS NULL);
INSERT INTO plugin_rehab_card ("category_id", "title", "label", "image_url", "source", "license", "difficulty", "status", "sort", "remark")
SELECT c."id", '脸', '脸', '/api/v1/rehab/card/image/wiki?file=Face_(Unsplash).jpg', 'wikimedia', 'Wikimedia Commons', 2, 1, 73, 'Real object photo via image proxy'
FROM plugin_rehab_card_category c WHERE c."code" = 'body-part'
  AND NOT EXISTS (SELECT 1 FROM plugin_rehab_card card WHERE card."category_id" = c."id" AND card."title" = '脸' AND card."deleted_at" IS NULL);
