// This file wires the crawler controller dependencies.

package crawler

import (
	crawlerapi "lina-plugin-rehab-cards/backend/api/crawler"
	v1 "lina-plugin-rehab-cards/backend/api/crawler/v1"
	crawlersvc "lina-plugin-rehab-cards/backend/internal/service/crawler"
)

// ControllerV1 is the picture-card crawler controller.
type ControllerV1 struct{ crawlerSvc crawlersvc.Service }

// NewV1 creates a picture-card crawler controller.
func NewV1() crawlerapi.ICrawlerV1 { return &ControllerV1{crawlerSvc: crawlersvc.New()} }

// toAPIEntity converts a service entity to an API projection.
func toAPIEntity(e *crawlersvc.Entity) *v1.JobEntity {
	if e == nil {
		return nil
	}
	return &v1.JobEntity{Id: e.Id, CategoryId: e.CategoryId, Keyword: e.Keyword, Provider: e.Provider, RequestedCount: e.RequestedCount, FetchedCount: e.FetchedCount, Status: e.Status, Message: e.Message, CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt}
}
