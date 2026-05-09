// This file implements the crawler run endpoint.

package crawler

import (
	"context"

	v1 "lina-plugin-rehab-cards/backend/api/crawler/v1"
	crawlersvc "lina-plugin-rehab-cards/backend/internal/service/crawler"
)

// Run creates a crawler job and placeholder cards.
func (c *ControllerV1) Run(ctx context.Context, req *v1.RunReq) (res *v1.RunRes, err error) {
	jobId, err := c.crawlerSvc.Run(ctx, crawlersvc.RunInput{CategoryId: req.CategoryId, Keyword: req.Keyword, Provider: req.Provider, Count: req.Count})
	if err != nil {
		return nil, err
	}
	return &v1.RunRes{JobId: jobId}, nil
}
