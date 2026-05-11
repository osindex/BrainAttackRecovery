// This file defines the card API contract implemented by the rehab-cards plugin.

package card

import (
	"context"

	"lina-plugin-rehab-cards/backend/api/card/v1"
)

// ICardV1 declares authenticated picture-card endpoints used by the admin UI.
type ICardV1 interface {
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}

// ICardPublicV1 declares public picture-card endpoints used by the patient H5.
type ICardPublicV1 interface {
	PublicList(ctx context.Context, req *v1.PublicListReq) (res *v1.PublicListRes, err error)
}
