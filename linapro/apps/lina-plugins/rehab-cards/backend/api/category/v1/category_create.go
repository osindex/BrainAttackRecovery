// This file declares the create-category endpoint DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// CreateReq defines the request for creating a picture-card category.
type CreateReq struct {
	g.Meta `path:"/rehab/card/category" method:"post" tags:"Rehab Card Categories" summary:"Create picture-card category" dc:"Create one picture-card category for object naming and cognitive training." permission:"rehab:card:category:add"`
	Name   string `json:"name" v:"required|max-length:100" dc:"Category display name" eg:"Animals"`
	Code   string `json:"code" v:"required|max-length:100" dc:"Stable category code, unique across categories" eg:"animal"`
	Sort   int    `json:"sort" dc:"Sort weight, smaller values are displayed earlier" eg:"10"`
	Status int    `json:"status" d:"1" dc:"Status: 1=enabled, 0=disabled" eg:"1"`
	Remark string `json:"remark" dc:"Optional remark" eg:"Common animals"`
}

// CreateRes defines the response for creating a picture-card category.
type CreateRes struct {
	Id int64 `json:"id" dc:"Created category ID" eg:"1"`
}
