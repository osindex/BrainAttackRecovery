// This file wires the summary controller dependencies.

package summary

import (
	summaryapi "lina-plugin-rehab-records/backend/api/summary"
	v1 "lina-plugin-rehab-records/backend/api/summary/v1"
	summarysvc "lina-plugin-rehab-records/backend/internal/service/summary"
)

// ControllerV1 is the training summary controller.
type ControllerV1 struct{ summarySvc summarysvc.Service }

// NewV1 creates a training summary controller.
func NewV1() summaryapi.ISummaryV1 { return &ControllerV1{summarySvc: summarysvc.New()} }

// toAPIItem converts a service summary item to an API projection.
func toAPIItem(item *summarysvc.DailyItem) *v1.DailyItem {
	if item == nil {
		return nil
	}
	return &v1.DailyItem{Date: item.Date, TrainingType: item.TrainingType, RecordCount: item.RecordCount, DurationSeconds: item.DurationSeconds, RepsCount: item.RepsCount, GazeCount: item.GazeCount, CorrectCount: item.CorrectCount}
}
