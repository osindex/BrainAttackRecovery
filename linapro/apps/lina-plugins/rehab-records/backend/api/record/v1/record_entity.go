// This file declares training record response entities.

package v1

import "github.com/gogf/gf/v2/os/gtime"

// RecordEntity describes one rehabilitation training record.
type RecordEntity struct {
	Id              int64       `json:"id" dc:"Record ID" eg:"1"`
	ClientRecordId  string      `json:"clientRecordId" dc:"Client-generated idempotency key" eg:"walk-1715248800000"`
	TrainingType    string      `json:"trainingType" dc:"Training type: walk, fist_raise, eye_gaze, card_game" eg:"walk"`
	OccurredOn      string      `json:"occurredOn" dc:"Training date in YYYY-MM-DD format" eg:"2026-05-09"`
	DurationSeconds int         `json:"durationSeconds" dc:"Duration in seconds for timer-based exercises" eg:"600"`
	SetsCount       int         `json:"setsCount" dc:"Set count for fist-raise exercise" eg:"2"`
	RepsCount       int         `json:"repsCount" dc:"Repetition count for fist-raise exercise" eg:"10"`
	GazeCount       int         `json:"gazeCount" dc:"Left-right gaze repetition count" eg:"20"`
	CardId          int64       `json:"cardId" dc:"Picture-card ID for card game records" eg:"1"`
	IsCorrect       int         `json:"isCorrect" dc:"Card game correctness: 1=correct, 0=incorrect" eg:"1"`
	ReactionMs      int         `json:"reactionMs" dc:"Card game reaction time in milliseconds" eg:"1500"`
	Payload         string      `json:"payload" dc:"Additional JSON payload as text" eg:"{}"`
	Remark          string      `json:"remark" dc:"Optional remark" eg:"Morning session"`
	CreatedAt       *gtime.Time `json:"createdAt" dc:"Creation time" eg:"2026-05-09 10:00:00"`
}
