// This file declares crawler job response entities.

package v1

import "github.com/gogf/gf/v2/os/gtime"

// JobEntity describes one picture crawler job.
type JobEntity struct {
	Id             int64       `json:"id" dc:"Crawler job ID" eg:"1"`
	CategoryId     int64       `json:"categoryId" dc:"Category ID" eg:"1"`
	Keyword        string      `json:"keyword" dc:"Search keyword used by the crawler" eg:"apple"`
	Provider       string      `json:"provider" dc:"Image provider: pixabay, wikimedia, or manual-placeholder" eg:"manual-placeholder"`
	RequestedCount int         `json:"requestedCount" dc:"Requested image count" eg:"10"`
	FetchedCount   int         `json:"fetchedCount" dc:"Fetched image count" eg:"10"`
	Status         string      `json:"status" dc:"Job status: pending, completed, failed" eg:"completed"`
	Message        string      `json:"message" dc:"Human-readable job message" eg:"Generated placeholder cards"`
	CreatedAt      *gtime.Time `json:"createdAt" dc:"Creation time" eg:"2026-05-09 10:00:00"`
	UpdatedAt      *gtime.Time `json:"updatedAt" dc:"Last update time" eg:"2026-05-09 10:00:00"`
}
