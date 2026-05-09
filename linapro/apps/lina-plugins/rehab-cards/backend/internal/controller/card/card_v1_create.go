// This file implements the card create endpoint.

package card

import (
	"context"

	v1 "lina-plugin-rehab-cards/backend/api/card/v1"
)

// Create creates a picture card.
func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := c.cardSvc.Create(ctx, saveInputFromCreate(req))
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: id}, nil
}
