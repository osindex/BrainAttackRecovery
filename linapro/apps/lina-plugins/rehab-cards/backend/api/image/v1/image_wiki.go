// This file declares the Wikimedia image proxy DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// WikiImageReq defines the request for proxying a Wikimedia Commons image by file name.
// The endpoint is intentionally public so the patient H5 can use it as an <img> source
// without attaching a JWT; abuse is bounded by the strict file-name allowlist regex.
type WikiImageReq struct {
	g.Meta `path:"/rehab/card/image/wiki" method:"get" tags:"Rehab Card Images" summary:"Proxy a Wikimedia Commons image" dc:"Stream a Wikimedia Commons image through the LinaPro server, so the patient H5 does not need direct access to commons.wikimedia.org. The file name must match a safe Special:FilePath name regex registered by the proxy. Public endpoint, no authentication required."`
	File   string `json:"file" v:"required|max-length:200" dc:"Wikimedia Commons file name, for example Red_Apple.jpg" eg:"Red_Apple.jpg"`
}

// WikiImageRes is intentionally empty because the controller streams the image bytes directly.
type WikiImageRes struct{}
