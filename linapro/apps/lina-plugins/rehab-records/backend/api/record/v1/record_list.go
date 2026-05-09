// This file declares the training record list endpoint DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// ListReq defines the request for listing training records.
type ListReq struct {
	g.Meta       `path:"/rehab/record" method:"get" tags:"Rehab Records" summary:"List training records" dc:"Paginated query for rehabilitation training records, supporting training type and date range filters." permission:"rehab:record:list"`
	PageNum      int    `json:"pageNum" d:"1" v:"min:1" dc:"Page number, starting from 1" eg:"1"`
	PageSize     int    `json:"pageSize" d:"20" v:"min:1|max:100" dc:"Number of items per page" eg:"20"`
	TrainingType string `json:"trainingType" dc:"Filter by training type: walk, fist_raise, eye_gaze, card_game, omitted means all types" eg:"walk"`
	StartDate    string `json:"startDate" v:"date-format:Y-m-d" dc:"Start date in YYYY-MM-DD format, omitted means no lower bound" eg:"2026-05-01"`
	EndDate      string `json:"endDate" v:"date-format:Y-m-d" dc:"End date in YYYY-MM-DD format, omitted means no upper bound" eg:"2026-05-31"`
}

// ListRes defines the response for listing training records.
type ListRes struct {
	List  []*RecordEntity `json:"list" dc:"Training record list" eg:"[]"`
	Total int             `json:"total" dc:"Total number of matched records" eg:"20"`
}
