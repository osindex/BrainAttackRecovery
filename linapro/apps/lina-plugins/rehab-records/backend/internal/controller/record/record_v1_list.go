// This file implements the training record list endpoint.

package record

import (
	"context"

	v1 "lina-plugin-rehab-records/backend/api/record/v1"
	recordsvc "lina-plugin-rehab-records/backend/internal/service/record"
)

// List queries training records.
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	out, err := c.recordSvc.List(ctx, recordsvc.ListInput{PageNum: req.PageNum, PageSize: req.PageSize, TrainingType: req.TrainingType, StartDate: req.StartDate, EndDate: req.EndDate})
	if err != nil {
		return nil, err
	}
	items := make([]*v1.RecordEntity, 0, len(out.List))
	for _, item := range out.List {
		items = append(items, toAPIEntity(item))
	}
	return &v1.ListRes{List: items, Total: out.Total}, nil
}
