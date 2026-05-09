// Package record implements rehabilitation training record persistence.
package record

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TableName stores the plugin-owned record table name.
const TableName = "plugin_rehab_record"

// Service defines the record service contract.
type Service interface {
	// List queries records with pagination and filters.
	List(ctx context.Context, in ListInput) (*ListOutput, error)
	// Create creates one training record and returns its ID.
	Create(ctx context.Context, in SaveInput) (int64, error)
	// Delete soft-deletes one training record by ID.
	Delete(ctx context.Context, id int64) error
}

// DuplicateRecordError reports that a client idempotency key already exists.
type DuplicateRecordError struct {
	// Id is the existing remote record ID.
	Id int64
}

// Error returns a stable diagnostic string for duplicate record inserts.
func (e *DuplicateRecordError) Error() string { return "duplicate client record id" }

// Ensure serviceImpl implements Service.
var _ Service = (*serviceImpl)(nil)

// serviceImpl implements Service.
type serviceImpl struct{}

// New creates and returns a record service instance.
func New() Service { return &serviceImpl{} }

// Entity describes one training record row.
type Entity struct {
	Id              int64       `orm:"id" json:"id"`
	ClientRecordId  string      `orm:"client_record_id" json:"clientRecordId"`
	TrainingType    string      `orm:"training_type" json:"trainingType"`
	OccurredOn      string      `orm:"occurred_on" json:"occurredOn"`
	DurationSeconds int         `orm:"duration_seconds" json:"durationSeconds"`
	SetsCount       int         `orm:"sets_count" json:"setsCount"`
	RepsCount       int         `orm:"reps_count" json:"repsCount"`
	GazeCount       int         `orm:"gaze_count" json:"gazeCount"`
	CardId          int64       `orm:"card_id" json:"cardId"`
	IsCorrect       int         `orm:"is_correct" json:"isCorrect"`
	ReactionMs      int         `orm:"reaction_ms" json:"reactionMs"`
	Payload         string      `orm:"payload" json:"payload"`
	Remark          string      `orm:"remark" json:"remark"`
	CreatedAt       *gtime.Time `orm:"created_at" json:"createdAt"`
}

// ListInput defines record list filters.
type ListInput struct {
	PageNum      int
	PageSize     int
	TrainingType string
	StartDate    string
	EndDate      string
}

// ListOutput defines record list output.
type ListOutput struct {
	List  []*Entity
	Total int
}

// SaveInput defines fields used to create records.
type SaveInput struct {
	ClientRecordId  string
	TrainingType    string
	OccurredOn      string
	DurationSeconds int
	SetsCount       int
	RepsCount       int
	GazeCount       int
	CardId          int64
	IsCorrect       int
	ReactionMs      int
	Payload         string
	Remark          string
}

// saveRow maps training record write fields to database columns.
type saveRow struct {
	ClientRecordId  string `orm:"client_record_id"`
	TrainingType    string `orm:"training_type"`
	OccurredOn      string `orm:"occurred_on"`
	DurationSeconds int    `orm:"duration_seconds"`
	SetsCount       int    `orm:"sets_count"`
	RepsCount       int    `orm:"reps_count"`
	GazeCount       int    `orm:"gaze_count"`
	CardId          int64  `orm:"card_id"`
	IsCorrect       int    `orm:"is_correct"`
	ReactionMs      int    `orm:"reaction_ms"`
	Payload         string `orm:"payload"`
	Remark          string `orm:"remark"`
}

// List queries records with pagination and filters.
func (s *serviceImpl) List(ctx context.Context, in ListInput) (*ListOutput, error) {
	model := g.DB().Model(TableName).Ctx(ctx)
	if in.TrainingType != "" {
		model = model.Where("training_type", in.TrainingType)
	}
	if in.StartDate != "" {
		model = model.WhereGTE("occurred_on", in.StartDate)
	}
	if in.EndDate != "" {
		model = model.WhereLTE("occurred_on", in.EndDate)
	}
	total, err := model.Count()
	if err != nil {
		return nil, err
	}
	list := make([]*Entity, 0)
	err = model.Page(in.PageNum, in.PageSize).OrderDesc("occurred_on").OrderDesc("id").Scan(&list)
	if err != nil {
		return nil, err
	}
	return &ListOutput{List: list, Total: total}, nil
}

// Create creates one training record and returns its ID.
func (s *serviceImpl) Create(ctx context.Context, in SaveInput) (int64, error) {
	payload := in.Payload
	if payload == "" {
		payload = "{}"
	}
	var existing *Entity
	if err := g.DB().Model(TableName).Ctx(ctx).Where("client_record_id", in.ClientRecordId).Scan(&existing); err != nil {
		return 0, err
	}
	if existing != nil {
		return existing.Id, nil
	}
	return g.DB().Model(TableName).Ctx(ctx).Data(saveRow{ClientRecordId: in.ClientRecordId, TrainingType: in.TrainingType, OccurredOn: in.OccurredOn, DurationSeconds: in.DurationSeconds, SetsCount: in.SetsCount, RepsCount: in.RepsCount, GazeCount: in.GazeCount, CardId: in.CardId, IsCorrect: in.IsCorrect, ReactionMs: in.ReactionMs, Payload: payload, Remark: in.Remark}).InsertAndGetId()
}

// Delete soft-deletes one training record by ID.
func (s *serviceImpl) Delete(ctx context.Context, id int64) error {
	_, err := g.DB().Model(TableName).Ctx(ctx).Where("id", id).Delete()
	return err
}
