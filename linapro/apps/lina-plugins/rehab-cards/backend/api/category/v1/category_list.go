// This file declares the category list endpoint DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// ListReq defines the request for listing picture-card categories.
type ListReq struct {
	g.Meta   `path:"/rehab/card/category" method:"get" tags:"Rehab Card Categories" summary:"List picture-card categories" dc:"Paginated query for picture-card categories, supporting fuzzy name filtering and status filtering." permission:"rehab:card:category:list"`
	PageNum  int    `json:"pageNum" d:"1" v:"min:1" dc:"Page number, starting from 1" eg:"1"`
	PageSize int    `json:"pageSize" d:"20" v:"min:1|max:100" dc:"Number of items per page" eg:"20"`
	Name     string `json:"name" dc:"Filter by category name using fuzzy matching, omitted means all names" eg:"animal"`
	Status   int    `json:"status" dc:"Filter by status: 1=enabled, 0=disabled, omitted means all statuses" eg:"1"`
}

// ListRes defines the response for listing picture-card categories.
type ListRes struct {
	List  []*CategoryEntity `json:"list" dc:"Picture-card category list" eg:"[]"`
	Total int               `json:"total" dc:"Total number of matched categories" eg:"6"`
}
