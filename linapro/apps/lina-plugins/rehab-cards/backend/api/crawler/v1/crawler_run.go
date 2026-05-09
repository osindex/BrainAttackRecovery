// This file declares the run-crawler endpoint DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// RunReq defines the request for running a picture crawler job.
type RunReq struct {
	g.Meta     `path:"/rehab/card/crawler" method:"post" tags:"Rehab Card Crawler" summary:"Run picture crawler job" dc:"Create a crawler job for one category and keyword. The MVP stores placeholder cards synchronously so the admin workflow is available before external image APIs are configured." permission:"rehab:card:crawler:run"`
	CategoryId int64  `json:"categoryId" v:"required|min:1" dc:"Target category ID" eg:"1"`
	Keyword    string `json:"keyword" v:"required|max-length:100" dc:"Search keyword" eg:"apple"`
	Provider   string `json:"provider" d:"manual-placeholder" dc:"Image provider: pixabay, wikimedia, or manual-placeholder" eg:"manual-placeholder"`
	Count      int    `json:"count" d:"5" v:"min:1|max:50" dc:"Requested image count" eg:"5"`
}

// RunRes defines the response for running a picture crawler job.
type RunRes struct {
	JobId int64 `json:"jobId" dc:"Created crawler job ID" eg:"1"`
}
