// Package card implements picture-card persistence for the rehab-cards plugin.
package card

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TableName stores the plugin-owned card table name.
const TableName = "plugin_rehab_card"

// Service defines the card service contract.
type Service interface {
	// List queries picture cards with pagination and filters.
	List(ctx context.Context, in ListInput) (*ListOutput, error)
	// Create creates one picture card and returns its generated ID.
	Create(ctx context.Context, in SaveInput) (int64, error)
	// Update updates one picture card by ID.
	Update(ctx context.Context, id int64, in SaveInput) error
	// Delete soft-deletes one picture card by ID.
	Delete(ctx context.Context, id int64) error
}

// Ensure serviceImpl implements Service.
var _ Service = (*serviceImpl)(nil)

// serviceImpl implements Service.
type serviceImpl struct{}

// New creates and returns a card service instance.
func New() Service { return &serviceImpl{} }

// Entity describes one picture-card row plus category projection.
type Entity struct {
	Id           int64       `orm:"id" json:"id"`
	CategoryId   int64       `orm:"category_id" json:"categoryId"`
	CategoryName string      `orm:"category_name" json:"categoryName"`
	Title        string      `orm:"title" json:"title"`
	Label        string      `orm:"label" json:"label"`
	ImageUrl     string      `orm:"image_url" json:"imageUrl"`
	ImageFileId  int64       `orm:"image_file_id" json:"imageFileId"`
	Source       string      `orm:"source" json:"source"`
	SourceUrl    string      `orm:"source_url" json:"sourceUrl"`
	License      string      `orm:"license" json:"license"`
	Difficulty   int         `orm:"difficulty" json:"difficulty"`
	Status       int         `orm:"status" json:"status"`
	Sort         int         `orm:"sort" json:"sort"`
	Remark       string      `orm:"remark" json:"remark"`
	CreatedAt    *gtime.Time `orm:"created_at" json:"createdAt"`
	UpdatedAt    *gtime.Time `orm:"updated_at" json:"updatedAt"`
}

// ListInput defines card list filters.
type ListInput struct {
	PageNum    int
	PageSize   int
	CategoryId int64
	Title      string
	Status     int
	Difficulty int
}

// ListOutput defines card list output.
type ListOutput struct {
	List  []*Entity
	Total int
}

// SaveInput defines fields used for create and update operations.
type SaveInput struct {
	CategoryId  int64
	Title       string
	Label       string
	ImageUrl    string
	ImageFileId int64
	Source      string
	SourceUrl   string
	License     string
	Difficulty  int
	Status      int
	Sort        int
	Remark      string
}

// saveRow maps picture-card write fields to database columns.
type saveRow struct {
	CategoryId  int64  `orm:"category_id"`
	Title       string `orm:"title"`
	Label       string `orm:"label"`
	ImageUrl    string `orm:"image_url"`
	ImageFileId int64  `orm:"image_file_id"`
	Source      string `orm:"source"`
	SourceUrl   string `orm:"source_url"`
	License     string `orm:"license"`
	Difficulty  int    `orm:"difficulty"`
	Status      int    `orm:"status"`
	Sort        int    `orm:"sort"`
	Remark      string `orm:"remark"`
}

// List queries picture cards with pagination and filters.
func (s *serviceImpl) List(ctx context.Context, in ListInput) (*ListOutput, error) {
	base := g.DB().Model(TableName+" c").Ctx(ctx).LeftJoin("plugin_rehab_card_category cat", "cat.id=c.category_id")
	if in.CategoryId > 0 {
		base = base.Where("c.category_id", in.CategoryId)
	}
	if in.Title != "" {
		base = base.WhereLike("c.title", "%"+in.Title+"%")
	}
	if in.Status == 0 || in.Status == 1 {
		base = base.Where("c.status", in.Status)
	}
	if in.Difficulty > 0 {
		base = base.Where("c.difficulty", in.Difficulty)
	}
	total, err := base.Clone().Fields("c.id").Count()
	if err != nil {
		return nil, err
	}
	list := make([]*Entity, 0)
	err = base.Fields("c.*, cat.name AS category_name").Page(in.PageNum, in.PageSize).OrderAsc("c.sort").OrderDesc("c.id").Scan(&list)
	if err != nil {
		return nil, err
	}
	return &ListOutput{List: list, Total: total}, nil
}

// Create creates one picture card and returns its generated ID.
func (s *serviceImpl) Create(ctx context.Context, in SaveInput) (int64, error) {
	return g.DB().Model(TableName).Ctx(ctx).Data(saveRow{CategoryId: in.CategoryId, Title: in.Title, Label: in.Label, ImageUrl: in.ImageUrl, ImageFileId: in.ImageFileId, Source: in.Source, SourceUrl: in.SourceUrl, License: in.License, Difficulty: in.Difficulty, Status: in.Status, Sort: in.Sort, Remark: in.Remark}).InsertAndGetId()
}

// Update updates one picture card by ID.
func (s *serviceImpl) Update(ctx context.Context, id int64, in SaveInput) error {
	_, err := g.DB().Model(TableName).Ctx(ctx).Where("id", id).Data(saveRow{CategoryId: in.CategoryId, Title: in.Title, Label: in.Label, ImageUrl: in.ImageUrl, ImageFileId: in.ImageFileId, Source: in.Source, SourceUrl: in.SourceUrl, License: in.License, Difficulty: in.Difficulty, Status: in.Status, Sort: in.Sort, Remark: in.Remark}).Update()
	return err
}

// Delete soft-deletes one picture card by ID.
func (s *serviceImpl) Delete(ctx context.Context, id int64) error {
	_, err := g.DB().Model(TableName).Ctx(ctx).Where("id", id).Delete()
	return err
}
