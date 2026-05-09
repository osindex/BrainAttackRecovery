// This file declares the delete-record endpoint DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// DeleteReq defines the request for deleting a training record.
type DeleteReq struct {
	g.Meta `path:"/rehab/record/{id}" method:"delete" tags:"Rehab Records" summary:"Delete training record" dc:"Soft-delete one rehabilitation training record by ID." permission:"rehab:record:remove"`
	Id     int64 `json:"id" v:"required|min:1" dc:"Record ID" eg:"1"`
}

// DeleteRes defines the response for deleting a training record.
type DeleteRes struct{}
