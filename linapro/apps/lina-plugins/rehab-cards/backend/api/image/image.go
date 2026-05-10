// This file defines the image proxy API contract for the rehab-cards plugin.

package image

import (
	"context"

	"lina-plugin-rehab-cards/backend/api/image/v1"
)

// IImageV1 declares image proxy endpoints.
type IImageV1 interface {
	WikiImage(ctx context.Context, req *v1.WikiImageReq) (res *v1.WikiImageRes, err error)
}
