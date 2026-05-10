// This file declares the Wikimedia image proxy DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// WikiImageReq defines the request for proxying a Wikimedia Commons image by file name.
type WikiImageReq struct {
	g.Meta `path:"/rehab/card/image/wiki" method:"get" tags:"Rehab Card Images" summary:"Proxy a Wikimedia Commons image" dc:"Stream a Wikimedia Commons image through the LinaPro server, so the patient H5 does not need direct access to commons.wikimedia.org. The file name must be a Special:FilePath name registered in the proxy allowlist or matching a safe character set." permission:"rehab:card:list"`
	File   string `json:"file" v:"required|max-length:200" dc:"Wikimedia Commons file name, for example Red_Apple.jpg" eg:"Red_Apple.jpg"`
}

// WikiImageRes is intentionally empty because the controller streams the image bytes directly.
type WikiImageRes struct{}
