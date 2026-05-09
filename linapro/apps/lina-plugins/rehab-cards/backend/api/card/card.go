// This file defines the card API contract implemented by the rehab-cards plugin.

package card

import (
	"context"

	"lina-plugin-rehab-cards/backend/api/card/v1"
)

// ICardV1 declares picture-card endpoints.
type ICardV1 interface {
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}
