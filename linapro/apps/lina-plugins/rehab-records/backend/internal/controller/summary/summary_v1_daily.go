// This file implements the daily summary endpoint.

package summary

import (
	"context"

	v1 "lina-plugin-rehab-records/backend/api/summary/v1"
	summarysvc "lina-plugin-rehab-records/backend/internal/service/summary"
)

// Daily returns chart-ready daily summaries.
func (c *ControllerV1) Daily(ctx context.Context, req *v1.DailyReq) (res *v1.DailyRes, err error) {
	items, err := c.summarySvc.Daily(ctx, summarysvc.DailyInput{TrainingType: req.TrainingType, StartDate: req.StartDate, EndDate: req.EndDate})
	if err != nil {
		return nil, err
	}
	list := make([]*v1.DailyItem, 0, len(items))
	for _, item := range items {
		list = append(list, toAPIItem(item))
	}
	return &v1.DailyRes{List: list}, nil
}
