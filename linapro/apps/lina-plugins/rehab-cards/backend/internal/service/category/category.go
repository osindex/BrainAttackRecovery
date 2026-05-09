// Package category implements picture-card category persistence for the rehab-cards plugin.
package category

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TableName stores the plugin-owned category table name.
const TableName = "plugin_rehab_card_category"

// Service defines the category service contract.
type Service interface {
	// List queries category records with pagination and filters.
	List(ctx context.Context, in ListInput) (*ListOutput, error)
	// Create creates one category and returns its generated ID.
	Create(ctx context.Context, in SaveInput) (int64, error)
	// Update updates one category by ID.
	Update(ctx context.Context, id int64, in SaveInput) error
	// Delete soft-deletes one category by ID.
	Delete(ctx context.Context, id int64) error
}

// Ensure serviceImpl implements Service.
var _ Service = (*serviceImpl)(nil)

// serviceImpl implements Service.
type serviceImpl struct{}

// New creates and returns a category service instance.
func New() Service { return &serviceImpl{} }

// Entity describes one category row.
type Entity struct {
	Id        int64       `orm:"id" json:"id"`
	Name      string      `orm:"name" json:"name"`
	Code      string      `orm:"code" json:"code"`
	Sort      int         `orm:"sort" json:"sort"`
	Status    int         `orm:"status" json:"status"`
	Remark    string      `orm:"remark" json:"remark"`
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"`
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"`
}

// ListInput defines category list filters.
type ListInput struct {
	PageNum  int
	PageSize int
	Name     string
	Status   int
}

// ListOutput defines category list output.
type ListOutput struct {
	List  []*Entity
	Total int
}

// SaveInput defines fields used for create and update operations.
type SaveInput struct {
	Name   string
	Code   string
	Sort   int
	Status int
	Remark string
}

// saveRow maps category write fields to database columns.
type saveRow struct {
	Name   string `orm:"name"`
	Code   string `orm:"code"`
	Sort   int    `orm:"sort"`
	Status int    `orm:"status"`
	Remark string `orm:"remark"`
}

// List queries category records with pagination and filters.
func (s *serviceImpl) List(ctx context.Context, in ListInput) (*ListOutput, error) {
	model := g.DB().Model(TableName).Ctx(ctx)
	if in.Name != "" {
		model = model.WhereLike("name", "%"+in.Name+"%")
	}
	if in.Status == 0 || in.Status == 1 {
		model = model.Where("status", in.Status)
	}
	total, err := model.Count()
	if err != nil {
		return nil, err
	}
	list := make([]*Entity, 0)
	err = model.Page(in.PageNum, in.PageSize).OrderAsc("sort").OrderAsc("id").Scan(&list)
	if err != nil {
		return nil, err
	}
	return &ListOutput{List: list, Total: total}, nil
}

// Create creates one category and returns its generated ID.
func (s *serviceImpl) Create(ctx context.Context, in SaveInput) (int64, error) {
	return g.DB().Model(TableName).Ctx(ctx).Data(saveRow{Name: in.Name, Code: in.Code, Sort: in.Sort, Status: in.Status, Remark: in.Remark}).InsertAndGetId()
}

// Update updates one category by ID.
func (s *serviceImpl) Update(ctx context.Context, id int64, in SaveInput) error {
	_, err := g.DB().Model(TableName).Ctx(ctx).Where("id", id).Data(saveRow{Name: in.Name, Code: in.Code, Sort: in.Sort, Status: in.Status, Remark: in.Remark}).Update()
	return err
}

// Delete soft-deletes one category by ID.
func (s *serviceImpl) Delete(ctx context.Context, id int64) error {
	_, err := g.DB().Model(TableName).Ctx(ctx).Where("id", id).Delete()
	return err
}
