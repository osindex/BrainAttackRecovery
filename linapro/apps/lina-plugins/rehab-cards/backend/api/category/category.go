// This file defines the category API contract implemented by the rehab-cards plugin.

package category

import (
	"context"

	"lina-plugin-rehab-cards/backend/api/category/v1"
)

// ICategoryV1 declares picture-card category endpoints.
type ICategoryV1 interface {
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}
