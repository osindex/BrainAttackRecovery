// This file declares the card list endpoint DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// ListReq defines the request for listing picture cards.
type ListReq struct {
	g.Meta     `path:"/rehab/card" method:"get" tags:"Rehab Cards" summary:"List picture cards" dc:"Paginated query for picture cards, supporting category, title, status, and difficulty filters." permission:"rehab:card:list"`
	PageNum    int    `json:"pageNum" d:"1" v:"min:1" dc:"Page number, starting from 1" eg:"1"`
	PageSize   int    `json:"pageSize" d:"20" v:"min:1|max:100" dc:"Number of items per page" eg:"20"`
	CategoryId int64  `json:"categoryId" dc:"Filter by category ID, omitted means all categories" eg:"1"`
	Title      string `json:"title" dc:"Filter by card title using fuzzy matching, omitted means all titles" eg:"apple"`
	Status     int    `json:"status" dc:"Filter by status: 1=enabled, 0=disabled, omitted means all statuses" eg:"1"`
	Difficulty int    `json:"difficulty" dc:"Filter by difficulty: 1=easy, 2=normal, 3=hard, omitted means all difficulties" eg:"1"`
}

// ListRes defines the response for listing picture cards.
type ListRes struct {
	List  []*CardEntity `json:"list" dc:"Picture-card list" eg:"[]"`
	Total int           `json:"total" dc:"Total number of matched picture cards" eg:"100"`
}
