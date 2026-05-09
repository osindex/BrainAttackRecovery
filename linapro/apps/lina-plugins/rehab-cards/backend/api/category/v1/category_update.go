// This file declares the update-category endpoint DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// UpdateReq defines the request for updating a picture-card category.
type UpdateReq struct {
	g.Meta `path:"/rehab/card/category/{id}" method:"put" tags:"Rehab Card Categories" summary:"Update picture-card category" dc:"Update one picture-card category by ID." permission:"rehab:card:category:edit"`
	Id     int64  `json:"id" v:"required|min:1" dc:"Category ID" eg:"1"`
	Name   string `json:"name" v:"required|max-length:100" dc:"Category display name" eg:"Animals"`
	Code   string `json:"code" v:"required|max-length:100" dc:"Stable category code, unique across categories" eg:"animal"`
	Sort   int    `json:"sort" dc:"Sort weight, smaller values are displayed earlier" eg:"10"`
	Status int    `json:"status" dc:"Status: 1=enabled, 0=disabled" eg:"1"`
	Remark string `json:"remark" dc:"Optional remark" eg:"Common animals"`
}

// UpdateRes defines the response for updating a picture-card category.
type UpdateRes struct{}
