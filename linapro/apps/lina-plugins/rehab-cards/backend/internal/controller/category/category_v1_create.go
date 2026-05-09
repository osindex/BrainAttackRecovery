// This file implements the category create endpoint.

package category

import (
	"context"

	v1 "lina-plugin-rehab-cards/backend/api/category/v1"
	categorysvc "lina-plugin-rehab-cards/backend/internal/service/category"
)

// Create creates a picture-card category.
func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := c.categorySvc.Create(ctx, categorysvc.SaveInput{Name: req.Name, Code: req.Code, Sort: req.Sort, Status: req.Status, Remark: req.Remark})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: id}, nil
}
