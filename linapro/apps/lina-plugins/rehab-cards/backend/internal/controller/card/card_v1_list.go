// This file implements the card list endpoint.

package card

import (
	"context"

	v1 "lina-plugin-rehab-cards/backend/api/card/v1"
	cardsvc "lina-plugin-rehab-cards/backend/internal/service/card"
)

// List queries picture cards.
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	out, err := c.cardSvc.List(ctx, cardsvc.ListInput{PageNum: req.PageNum, PageSize: req.PageSize, CategoryId: req.CategoryId, Title: req.Title, Status: req.Status, Difficulty: req.Difficulty})
	if err != nil {
		return nil, err
	}
	items := make([]*v1.CardEntity, 0, len(out.List))
	for _, item := range out.List {
		items = append(items, toAPIEntity(item))
	}
	return &v1.ListRes{List: items, Total: out.Total}, nil
}
