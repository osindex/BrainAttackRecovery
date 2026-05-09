// This file implements the card update endpoint.

package card

import (
	"context"

	v1 "lina-plugin-rehab-cards/backend/api/card/v1"
)

// Update updates a picture card.
func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	if err = c.cardSvc.Update(ctx, req.Id, saveInputFromUpdate(req)); err != nil {
		return nil, err
	}
	return &v1.UpdateRes{}, nil
}
