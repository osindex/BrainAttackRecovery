// This file implements the card delete endpoint.

package card

import (
	"context"

	v1 "lina-plugin-rehab-cards/backend/api/card/v1"
)

// Delete soft-deletes a picture card.
func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	if err = c.cardSvc.Delete(ctx, req.Id); err != nil {
		return nil, err
	}
	return &v1.DeleteRes{}, nil
}
