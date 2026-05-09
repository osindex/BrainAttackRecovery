// This file implements the training record delete endpoint.

package record

import (
	"context"

	v1 "lina-plugin-rehab-records/backend/api/record/v1"
)

// Delete soft-deletes a training record.
func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	if err = c.recordSvc.Delete(ctx, req.Id); err != nil {
		return nil, err
	}
	return &v1.DeleteRes{}, nil
}
