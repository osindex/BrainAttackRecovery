// This file implements the training record create endpoint.

package record

import (
	"context"

	v1 "lina-plugin-rehab-records/backend/api/record/v1"
)

// Create creates a training record.
func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := c.recordSvc.Create(ctx, saveInputFromCreate(req))
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: id}, nil
}
