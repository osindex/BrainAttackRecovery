// Package seed maintains the canonical rehab-card catalog and synchronizes it
// into the database. It runs once at host startup so a fresh deployment always
// has playable cards even before any operator action, and on a daily cron so
// later catalog additions reach the patient H5 without manual intervention.

package seed

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"lina-core/pkg/logger"
)

// Service defines the seeding contract for the rehab-card catalog.
type Service interface {
	// Sync inserts any missing canonical cards. Existing rows are left untouched
	// so manual administrator edits and deletions are preserved.
	Sync(ctx context.Context) (inserted int, err error)
}

// Ensure serviceImpl implements Service.
var _ Service = (*serviceImpl)(nil)

// serviceImpl implements Service.
type serviceImpl struct{}

// New creates and returns a seeding service instance.
func New() Service { return &serviceImpl{} }

// catalogCard is one immutable catalog entry expressed in Go so deployments do
// not need to reload SQL fixtures to refresh the picture-card library.
type catalogCard struct {
	categoryCode string
	title        string
	label        string
	imageURL     string
	difficulty   int
	sort         int
}

// categorySeed describes one canonical card category.
type categorySeed struct {
	code   string
	name   string
	sort   int
	remark string
}

// canonicalCategories lists the rehab-card categories ensured at startup.
var canonicalCategories = []categorySeed{
	{code: "animal", name: "动物", sort: 10, remark: "常见动物命名训练"},
	{code: "food", name: "食物", sort: 20, remark: "常见食物命名训练"},
	{code: "daily-object", name: "日用品", sort: 30, remark: "日常用品命名训练"},
	{code: "vehicle", name: "交通工具", sort: 40, remark: "交通工具命名训练"},
	{code: "clothing", name: "衣物", sort: 50, remark: "衣物命名训练"},
	{code: "body-part", name: "身体部位", sort: 60, remark: "身体部位命名训练"},
}

// canonicalCards lists the rehab-card library ensured at startup. Image URLs
// point directly to upload.wikimedia.org so the patient H5 fetches them without
// any backend proxy.
var canonicalCards = []catalogCard{
	// food
	{categoryCode: "food", title: "苹果", label: "苹果", imageURL: "https://upload.wikimedia.org/wikipedia/commons/1/15/Red_Apple.jpg", difficulty: 1, sort: 10},
	{categoryCode: "food", title: "香蕉", label: "香蕉", imageURL: "https://upload.wikimedia.org/wikipedia/commons/d/de/Bananavarieties.jpg", difficulty: 1, sort: 11},
	{categoryCode: "food", title: "橙子", label: "橙子", imageURL: "https://upload.wikimedia.org/wikipedia/commons/8/85/Orange_fruit.jpg", difficulty: 1, sort: 12},
	{categoryCode: "food", title: "西瓜", label: "西瓜", imageURL: "https://upload.wikimedia.org/wikipedia/commons/a/ae/Watermelon_cross_BNC.jpg", difficulty: 1, sort: 13},
	{categoryCode: "food", title: "草莓", label: "草莓", imageURL: "https://upload.wikimedia.org/wikipedia/commons/f/fc/Strawberry_Single1.jpg", difficulty: 2, sort: 14},
	{categoryCode: "food", title: "梨", label: "梨", imageURL: "https://upload.wikimedia.org/wikipedia/commons/d/dd/Yellow_pear_on_a_black_background.jpg", difficulty: 2, sort: 15},
	// daily-object
	{categoryCode: "daily-object", title: "杯子", label: "杯子", imageURL: "https://upload.wikimedia.org/wikipedia/commons/f/fb/White_cup_and_saucer.jpg", difficulty: 1, sort: 20},
	{categoryCode: "daily-object", title: "勺子", label: "勺子", imageURL: "https://upload.wikimedia.org/wikipedia/commons/7/70/Spoon_silver.jpg", difficulty: 2, sort: 21},
	{categoryCode: "daily-object", title: "筷子", label: "筷子", imageURL: "https://upload.wikimedia.org/wikipedia/commons/e/ea/Chopsticks-candc.jpg", difficulty: 2, sort: 22},
	{categoryCode: "daily-object", title: "水壶", label: "水壶", imageURL: "https://upload.wikimedia.org/wikipedia/commons/c/c8/WWII_Allied_Canteen.jpg", difficulty: 3, sort: 23},
	// animal
	{categoryCode: "animal", title: "小狗", label: "小狗", imageURL: "https://upload.wikimedia.org/wikipedia/commons/9/93/Golden_Retriever_Carlos_%2810581910556%29.jpg", difficulty: 1, sort: 30},
	{categoryCode: "animal", title: "小猫", label: "小猫", imageURL: "https://upload.wikimedia.org/wikipedia/commons/b/b6/Felis_catus-cat_on_snow.jpg", difficulty: 1, sort: 31},
	{categoryCode: "animal", title: "兔子", label: "兔子", imageURL: "https://upload.wikimedia.org/wikipedia/commons/1/1f/Oryctolagus_cuniculus_Rcdo.jpg", difficulty: 2, sort: 32},
	{categoryCode: "animal", title: "马", label: "马", imageURL: "https://upload.wikimedia.org/wikipedia/commons/0/01/Hauspferd.JPG", difficulty: 2, sort: 33},
	{categoryCode: "animal", title: "鸡", label: "鸡", imageURL: "https://upload.wikimedia.org/wikipedia/commons/5/5e/Gallus_gallus_domesticus_Brown_Leghorn.jpg", difficulty: 2, sort: 34},
	{categoryCode: "animal", title: "牛", label: "牛", imageURL: "https://upload.wikimedia.org/wikipedia/commons/0/0c/Cow_female_black_white.jpg", difficulty: 2, sort: 35},
	{categoryCode: "animal", title: "羊", label: "羊", imageURL: "https://upload.wikimedia.org/wikipedia/commons/9/99/Sheep_in_field.JPG", difficulty: 2, sort: 36},
	{categoryCode: "animal", title: "金鱼", label: "金鱼", imageURL: "https://upload.wikimedia.org/wikipedia/commons/e/e9/Goldfish3.jpg", difficulty: 3, sort: 37},
	{categoryCode: "animal", title: "蝴蝶", label: "蝴蝶", imageURL: "https://upload.wikimedia.org/wikipedia/commons/6/63/Monarch_In_May.jpg", difficulty: 3, sort: 38},
	{categoryCode: "animal", title: "蚱蜢", label: "蚱蜢", imageURL: "https://upload.wikimedia.org/wikipedia/commons/b/b5/Grasshopper_1.jpg", difficulty: 3, sort: 39},
	{categoryCode: "animal", title: "熊猫", label: "熊猫", imageURL: "https://upload.wikimedia.org/wikipedia/commons/c/cd/Panda_Cub_from_Wolong%2C_Sichuan%2C_China.JPG", difficulty: 2, sort: 40},
	// vehicle
	{categoryCode: "vehicle", title: "汽车", label: "汽车", imageURL: "https://upload.wikimedia.org/wikipedia/commons/a/a4/2019_Toyota_Corolla_Icon_Tech_VVT-i_Hybrid_1.8.jpg", difficulty: 1, sort: 50},
	{categoryCode: "vehicle", title: "自行车", label: "自行车", imageURL: "https://upload.wikimedia.org/wikipedia/commons/c/cb/Bicycle_Abbey_Sprotbrough.jpg", difficulty: 2, sort: 51},
	{categoryCode: "vehicle", title: "公交车", label: "公交车", imageURL: "https://upload.wikimedia.org/wikipedia/commons/0/07/Bus_in_Hong_Kong.jpg", difficulty: 2, sort: 52},
	{categoryCode: "vehicle", title: "飞机", label: "飞机", imageURL: "https://upload.wikimedia.org/wikipedia/commons/8/82/Airbus_A380_blue_sky.jpg", difficulty: 2, sort: 53},
	{categoryCode: "vehicle", title: "救护车", label: "救护车", imageURL: "https://upload.wikimedia.org/wikipedia/commons/7/78/Ambulance_Berlin.jpg", difficulty: 3, sort: 54},
	// clothing
	{categoryCode: "clothing", title: "鞋子", label: "鞋子", imageURL: "https://upload.wikimedia.org/wikipedia/commons/3/35/Shoes_-_Nike_Air_Jordan_1_Retro_Banned_2016_-_sneakers.jpg", difficulty: 2, sort: 60},
	{categoryCode: "clothing", title: "裤子", label: "裤子", imageURL: "https://upload.wikimedia.org/wikipedia/commons/6/6e/Blue_Denim_Jeans.jpg", difficulty: 2, sort: 61},
	{categoryCode: "clothing", title: "围巾", label: "围巾", imageURL: "https://upload.wikimedia.org/wikipedia/commons/e/ec/Red_scarf.jpg", difficulty: 2, sort: 62},
	{categoryCode: "clothing", title: "帽子", label: "帽子", imageURL: "https://upload.wikimedia.org/wikipedia/commons/9/98/Winter_hat.jpg", difficulty: 2, sort: 63},
	{categoryCode: "clothing", title: "手套", label: "手套", imageURL: "https://upload.wikimedia.org/wikipedia/commons/5/59/Glove.jpg", difficulty: 3, sort: 64},
	// body-part
	{categoryCode: "body-part", title: "手", label: "手", imageURL: "https://upload.wikimedia.org/wikipedia/commons/6/66/Hand_%28sculpture%29.jpg", difficulty: 1, sort: 70},
	{categoryCode: "body-part", title: "眼睛", label: "眼睛", imageURL: "https://upload.wikimedia.org/wikipedia/commons/0/0a/Human_eye.jpg", difficulty: 1, sort: 71},
	{categoryCode: "body-part", title: "脚", label: "脚", imageURL: "https://upload.wikimedia.org/wikipedia/commons/b/ba/Pair_of_feet.jpg", difficulty: 2, sort: 72},
	{categoryCode: "body-part", title: "脸", label: "脸", imageURL: "https://upload.wikimedia.org/wikipedia/commons/7/70/Face_%28Unsplash%29.jpg", difficulty: 2, sort: 73},
}

// categoryRow maps canonical category inserts to database columns.
type categoryRow struct {
	Name   string `orm:"name"`
	Code   string `orm:"code"`
	Sort   int    `orm:"sort"`
	Status int    `orm:"status"`
	Remark string `orm:"remark"`
}

// cardRow maps canonical card inserts to database columns.
type cardRow struct {
	CategoryId int64  `orm:"category_id"`
	Title      string `orm:"title"`
	Label      string `orm:"label"`
	ImageUrl   string `orm:"image_url"`
	Source     string `orm:"source"`
	License    string `orm:"license"`
	Difficulty int    `orm:"difficulty"`
	Status     int    `orm:"status"`
	Sort       int    `orm:"sort"`
	Remark     string `orm:"remark"`
}

// Sync inserts any missing canonical categories and cards.
func (s *serviceImpl) Sync(ctx context.Context) (int, error) {
	categoryIDs, err := s.ensureCategories(ctx)
	if err != nil {
		return 0, err
	}
	return s.ensureCards(ctx, categoryIDs)
}

// ensureCategories inserts missing canonical categories and returns the
// category-code → ID map used to attach cards.
func (s *serviceImpl) ensureCategories(ctx context.Context) (map[string]int64, error) {
	ids := make(map[string]int64, len(canonicalCategories))
	for _, seed := range canonicalCategories {
		var existingID int64
		existingValue, err := g.DB().Model("plugin_rehab_card_category").Ctx(ctx).Where("code", seed.code).Fields("id").Value()
		if err != nil {
			return nil, err
		}
		existingID = existingValue.Int64()
		if existingID > 0 {
			ids[seed.code] = existingID
			continue
		}
		insertedID, err := g.DB().Model("plugin_rehab_card_category").Ctx(ctx).Data(categoryRow{
			Name:   seed.name,
			Code:   seed.code,
			Sort:   seed.sort,
			Status: 1,
			Remark: seed.remark,
		}).InsertAndGetId()
		if err != nil {
			return nil, err
		}
		ids[seed.code] = insertedID
	}
	return ids, nil
}

// ensureCards inserts any canonical card that is not already present, keyed by
// (category_id, title). Existing rows are left untouched so administrator
// edits, custom cards, and intentional deletions remain stable across restarts.
func (s *serviceImpl) ensureCards(ctx context.Context, categoryIDs map[string]int64) (int, error) {
	inserted := 0
	for _, card := range canonicalCards {
		categoryID, ok := categoryIDs[card.categoryCode]
		if !ok {
			logger.Warningf(ctx, "rehab-cards seed missing category code=%s for card=%s", card.categoryCode, card.title)
			continue
		}
		count, err := g.DB().Model("plugin_rehab_card").Ctx(ctx).Where("category_id", categoryID).Where("title", card.title).Count()
		if err != nil {
			return inserted, err
		}
		if count > 0 {
			continue
		}
		_, err = g.DB().Model("plugin_rehab_card").Ctx(ctx).Data(cardRow{
			CategoryId: categoryID,
			Title:      card.title,
			Label:      card.label,
			ImageUrl:   card.imageURL,
			Source:     "wikimedia",
			License:    "Wikimedia Commons",
			Difficulty: card.difficulty,
			Status:     1,
			Sort:       card.sort,
			Remark:     "Seeded by rehab-cards catalog",
		}).Insert()
		if err != nil {
			return inserted, err
		}
		inserted++
	}
	return inserted, nil
}
