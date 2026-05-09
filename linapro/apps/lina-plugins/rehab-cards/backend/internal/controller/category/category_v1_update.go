// This file implements the category update endpoint.

package category

import (
	"context"

	v1 "lina-plugin-rehab-cards/backend/api/category/v1"
	categorysvc "lina-plugin-rehab-cards/backend/internal/service/category"
)

// Update updates a picture-card category.
func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = c.categorySvc.Update(ctx, req.Id, categorysvc.SaveInput{Name: req.Name, Code: req.Code, Sort: req.Sort, Status: req.Status, Remark: req.Remark})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateRes{}, nil
}
