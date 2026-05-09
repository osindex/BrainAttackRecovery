// Package summary implements chart-ready training record aggregations.
package summary

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// Service defines the summary service contract.
type Service interface {
	// Daily returns daily aggregates grouped by date and training type.
	Daily(ctx context.Context, in DailyInput) ([]*DailyItem, error)
}

// Ensure serviceImpl implements Service.
var _ Service = (*serviceImpl)(nil)

// serviceImpl implements Service.
type serviceImpl struct{}

// New creates and returns a summary service instance.
func New() Service { return &serviceImpl{} }

// DailyInput defines daily summary filters.
type DailyInput struct {
	TrainingType string
	StartDate    string
	EndDate      string
}

// DailyItem describes one aggregated date bucket.
type DailyItem struct {
	Date            string `orm:"date" json:"date"`
	TrainingType    string `orm:"training_type" json:"trainingType"`
	RecordCount     int    `orm:"record_count" json:"recordCount"`
	DurationSeconds int    `orm:"duration_seconds" json:"durationSeconds"`
	RepsCount       int    `orm:"reps_count" json:"repsCount"`
	GazeCount       int    `orm:"gaze_count" json:"gazeCount"`
	CorrectCount    int    `orm:"correct_count" json:"correctCount"`
}

// Daily returns daily aggregates grouped by date and training type.
func (s *serviceImpl) Daily(ctx context.Context, in DailyInput) ([]*DailyItem, error) {
	model := g.DB().Model("plugin_rehab_record").Ctx(ctx).Fields("occurred_on AS date, training_type, COUNT(*) AS record_count, SUM(duration_seconds) AS duration_seconds, SUM(reps_count) AS reps_count, SUM(gaze_count) AS gaze_count, SUM(is_correct) AS correct_count")
	if in.TrainingType != "" {
		model = model.Where("training_type", in.TrainingType)
	}
	if in.StartDate != "" {
		model = model.WhereGTE("occurred_on", in.StartDate)
	}
	if in.EndDate != "" {
		model = model.WhereLTE("occurred_on", in.EndDate)
	}
	items := make([]*DailyItem, 0)
	err := model.Group("occurred_on, training_type").OrderAsc("occurred_on").Scan(&items)
	return items, err
}
