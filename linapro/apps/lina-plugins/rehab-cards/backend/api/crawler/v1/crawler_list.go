// This file declares the crawler job list endpoint DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// ListReq defines the request for listing picture crawler jobs.
type ListReq struct {
	g.Meta     `path:"/rehab/card/crawler" method:"get" tags:"Rehab Card Crawler" summary:"List picture crawler jobs" dc:"Paginated query for picture-card crawler jobs." permission:"rehab:card:crawler:list"`
	PageNum    int    `json:"pageNum" d:"1" v:"min:1" dc:"Page number, starting from 1" eg:"1"`
	PageSize   int    `json:"pageSize" d:"20" v:"min:1|max:100" dc:"Number of items per page" eg:"20"`
	CategoryId int64  `json:"categoryId" dc:"Filter by category ID, omitted means all categories" eg:"1"`
	Status     string `json:"status" dc:"Filter by job status, omitted means all statuses" eg:"completed"`
}

// ListRes defines the response for listing picture crawler jobs.
type ListRes struct {
	List  []*JobEntity `json:"list" dc:"Crawler job list" eg:"[]"`
	Total int          `json:"total" dc:"Total number of matched crawler jobs" eg:"3"`
}
