// This file implements the crawler job list endpoint.

package crawler

import (
	"context"

	v1 "lina-plugin-rehab-cards/backend/api/crawler/v1"
	crawlersvc "lina-plugin-rehab-cards/backend/internal/service/crawler"
)

// List queries picture crawler jobs.
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	out, err := c.crawlerSvc.List(ctx, crawlersvc.ListInput{PageNum: req.PageNum, PageSize: req.PageSize, CategoryId: req.CategoryId, Status: req.Status})
	if err != nil {
		return nil, err
	}
	items := make([]*v1.JobEntity, 0, len(out.List))
	for _, item := range out.List {
		items = append(items, toAPIEntity(item))
	}
	return &v1.ListRes{List: items, Total: out.Total}, nil
}
