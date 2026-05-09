// This file implements the category list endpoint.

package category

import (
	"context"

	v1 "lina-plugin-rehab-cards/backend/api/category/v1"
	categorysvc "lina-plugin-rehab-cards/backend/internal/service/category"
)

// List queries picture-card categories.
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	out, err := c.categorySvc.List(ctx, categorysvc.ListInput{PageNum: req.PageNum, PageSize: req.PageSize, Name: req.Name, Status: req.Status})
	if err != nil {
		return nil, err
	}
	items := make([]*v1.CategoryEntity, 0, len(out.List))
	for _, item := range out.List {
		items = append(items, toAPIEntity(item))
	}
	return &v1.ListRes{List: items, Total: out.Total}, nil
}
