// This file implements the category delete endpoint.

package category

import (
	"context"

	v1 "lina-plugin-rehab-cards/backend/api/category/v1"
)

// Delete soft-deletes a picture-card category.
func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	if err = c.categorySvc.Delete(ctx, req.Id); err != nil {
		return nil, err
	}
	return &v1.DeleteRes{}, nil
}
