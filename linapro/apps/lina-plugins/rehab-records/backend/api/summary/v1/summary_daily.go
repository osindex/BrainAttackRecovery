// This file declares the daily training summary endpoint DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// DailyReq defines the request for daily training summary charts.
type DailyReq struct {
	g.Meta       `path:"/rehab/record/summary/daily" method:"get" tags:"Rehab Record Summary" summary:"Get daily training summary" dc:"Aggregate rehabilitation training records by date for chart rendering in the admin and H5 history pages." permission:"rehab:record:summary"`
	TrainingType string `json:"trainingType" dc:"Filter by training type: walk, fist_raise, eye_gaze, card_game, omitted means all types" eg:"walk"`
	StartDate    string `json:"startDate" v:"date-format:Y-m-d" dc:"Start date in YYYY-MM-DD format, omitted means no lower bound" eg:"2026-05-01"`
	EndDate      string `json:"endDate" v:"date-format:Y-m-d" dc:"End date in YYYY-MM-DD format, omitted means no upper bound" eg:"2026-05-31"`
}

// DailyItem describes one date bucket in the daily summary.
type DailyItem struct {
	Date            string `json:"date" dc:"Date in YYYY-MM-DD format" eg:"2026-05-09"`
	TrainingType    string `json:"trainingType" dc:"Training type: walk, fist_raise, eye_gaze, card_game" eg:"walk"`
	RecordCount     int    `json:"recordCount" dc:"Number of records in the bucket" eg:"3"`
	DurationSeconds int    `json:"durationSeconds" dc:"Total duration seconds" eg:"1200"`
	RepsCount       int    `json:"repsCount" dc:"Total repetition count" eg:"30"`
	GazeCount       int    `json:"gazeCount" dc:"Total gaze repetition count" eg:"20"`
	CorrectCount    int    `json:"correctCount" dc:"Correct card-game answer count" eg:"8"`
}

// DailyRes defines the response for daily training summaries.
type DailyRes struct {
	List []*DailyItem `json:"list" dc:"Daily training summary list" eg:"[]"`
}
