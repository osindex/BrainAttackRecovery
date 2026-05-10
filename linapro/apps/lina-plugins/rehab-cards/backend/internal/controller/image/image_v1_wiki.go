// This file implements the Wikimedia image proxy endpoint.

package image

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "lina-plugin-rehab-cards/backend/api/image/v1"
)

// WikiImage streams a Wikimedia Commons image through the LinaPro server.
func (c *ControllerV1) WikiImage(ctx context.Context, req *v1.WikiImageReq) (res *v1.WikiImageRes, err error) {
	contentType, body, err := c.imageSvc.FetchWikimedia(ctx, req.File)
	if err != nil {
		return nil, err
	}
	request := g.RequestFromCtx(ctx)
	if request == nil {
		return nil, nil
	}
	response := request.Response
	response.Header().Set("Content-Type", contentType)
	response.Header().Set("Cache-Control", "public, max-age=86400")
	response.Header().Set("X-Image-Source", "wikimedia-commons")
	response.Write(body)
	return nil, nil
}
