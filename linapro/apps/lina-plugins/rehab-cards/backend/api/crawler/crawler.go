// This file defines crawler API contracts implemented by the rehab-cards plugin.

package crawler

import (
	"context"

	"lina-plugin-rehab-cards/backend/api/crawler/v1"
)

// ICrawlerV1 declares picture-card crawler endpoints.
type ICrawlerV1 interface {
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Run(ctx context.Context, req *v1.RunReq) (res *v1.RunRes, err error)
}
