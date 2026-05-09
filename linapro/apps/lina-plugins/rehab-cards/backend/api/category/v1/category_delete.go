// This file declares the delete-category endpoint DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// DeleteReq defines the request for deleting a picture-card category.
type DeleteReq struct {
	g.Meta `path:"/rehab/card/category/{id}" method:"delete" tags:"Rehab Card Categories" summary:"Delete picture-card category" dc:"Soft-delete one picture-card category by ID." permission:"rehab:card:category:remove"`
	Id     int64 `json:"id" v:"required|min:1" dc:"Category ID" eg:"1"`
}

// DeleteRes defines the response for deleting a picture-card category.
type DeleteRes struct{}
