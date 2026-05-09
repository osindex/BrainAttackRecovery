// This file declares the create-record endpoint DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// CreateReq defines the request for creating a training record.
type CreateReq struct {
	g.Meta          `path:"/rehab/record" method:"post" tags:"Rehab Records" summary:"Create training record" dc:"Create one rehabilitation training record from the patient H5 app. clientRecordId is used for idempotent offline sync." permission:"rehab:record:add"`
	ClientRecordId  string `json:"clientRecordId" v:"required|max-length:100" dc:"Client-generated idempotency key" eg:"walk-1715248800000"`
	TrainingType    string `json:"trainingType" v:"required|in:walk,fist_raise,eye_gaze,card_game" dc:"Training type: walk, fist_raise, eye_gaze, card_game" eg:"walk"`
	OccurredOn      string `json:"occurredOn" v:"required|date-format:Y-m-d" dc:"Training date in YYYY-MM-DD format" eg:"2026-05-09"`
	DurationSeconds int    `json:"durationSeconds" dc:"Duration in seconds for timer-based exercises" eg:"600"`
	SetsCount       int    `json:"setsCount" dc:"Set count for fist-raise exercise" eg:"2"`
	RepsCount       int    `json:"repsCount" dc:"Repetition count for fist-raise exercise" eg:"10"`
	GazeCount       int    `json:"gazeCount" dc:"Left-right gaze repetition count" eg:"20"`
	CardId          int64  `json:"cardId" dc:"Picture-card ID for card game records" eg:"1"`
	IsCorrect       int    `json:"isCorrect" dc:"Card game correctness: 1=correct, 0=incorrect" eg:"1"`
	ReactionMs      int    `json:"reactionMs" dc:"Card game reaction time in milliseconds" eg:"1500"`
	Payload         string `json:"payload" v:"max-length:4000" dc:"Additional JSON payload as text" eg:"{}"`
	Remark          string `json:"remark" v:"max-length:500" dc:"Optional remark" eg:"Morning session"`
}

// CreateRes defines the response for creating a training record.
type CreateRes struct {
	Id int64 `json:"id" dc:"Created record ID" eg:"1"`
}
