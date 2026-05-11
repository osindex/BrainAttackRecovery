// This file implements the public picture-card list endpoint for the patient H5.

package card

import (
	"context"

	cardapi "lina-plugin-rehab-cards/backend/api/card"
	v1 "lina-plugin-rehab-cards/backend/api/card/v1"
	cardsvc "lina-plugin-rehab-cards/backend/internal/service/card"
)

// ControllerPublicV1 is the public picture-card controller exposed to the
// patient H5 without authentication. It returns enabled picture cards and lets
// the H5 render images directly from the upstream image URLs.
type ControllerPublicV1 struct{ cardSvc cardsvc.Service }

// NewPublicV1 creates the public picture-card controller.
func NewPublicV1() cardapi.ICardPublicV1 {
	return &ControllerPublicV1{cardSvc: cardsvc.New()}
}

// PublicList returns enabled picture cards for the patient H5.
func (c *ControllerPublicV1) PublicList(ctx context.Context, req *v1.PublicListReq) (res *v1.PublicListRes, err error) {
	limit := req.Limit
	if limit <= 0 {
		limit = 100
	}
	out, err := c.cardSvc.List(ctx, cardsvc.ListInput{
		PageNum:    1,
		PageSize:   limit,
		CategoryId: req.CategoryId,
		Title:      req.Keyword,
		Status:     1,
		Difficulty: req.Difficulty,
	})
	if err != nil {
		return nil, err
	}
	items := make([]*v1.CardEntity, 0, len(out.List))
	for _, item := range out.List {
		items = append(items, toAPIEntity(item))
	}
	return &v1.PublicListRes{List: items, Total: out.Total}, nil
}
