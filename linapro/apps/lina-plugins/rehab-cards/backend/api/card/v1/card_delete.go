// This file declares the delete-card endpoint DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// DeleteReq defines the request for deleting a picture card.
type DeleteReq struct {
	g.Meta `path:"/rehab/card/{id}" method:"delete" tags:"Rehab Cards" summary:"Delete picture card" dc:"Soft-delete one picture card by ID." permission:"rehab:card:remove"`
	Id     int64 `json:"id" v:"required|min:1" dc:"Card ID" eg:"1"`
}

// DeleteRes defines the response for deleting a picture card.
type DeleteRes struct{}
