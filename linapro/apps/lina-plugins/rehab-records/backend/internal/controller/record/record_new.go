// This file wires the training record controller dependencies.

package record

import (
	recordapi "lina-plugin-rehab-records/backend/api/record"
	v1 "lina-plugin-rehab-records/backend/api/record/v1"
	recordsvc "lina-plugin-rehab-records/backend/internal/service/record"
)

// ControllerV1 is the training record controller.
type ControllerV1 struct{ recordSvc recordsvc.Service }

// NewV1 creates a training record controller.
func NewV1() recordapi.IRecordV1 { return &ControllerV1{recordSvc: recordsvc.New()} }

// toAPIEntity converts a service entity to an API projection.
func toAPIEntity(e *recordsvc.Entity) *v1.RecordEntity {
	if e == nil {
		return nil
	}
	return &v1.RecordEntity{Id: e.Id, ClientRecordId: e.ClientRecordId, TrainingType: e.TrainingType, OccurredOn: e.OccurredOn, DurationSeconds: e.DurationSeconds, SetsCount: e.SetsCount, RepsCount: e.RepsCount, GazeCount: e.GazeCount, CardId: e.CardId, IsCorrect: e.IsCorrect, ReactionMs: e.ReactionMs, Payload: e.Payload, Remark: e.Remark, CreatedAt: e.CreatedAt}
}

// saveInputFromCreate converts a create request to service input.
func saveInputFromCreate(req *v1.CreateReq) recordsvc.SaveInput {
	return recordsvc.SaveInput{ClientRecordId: req.ClientRecordId, TrainingType: req.TrainingType, OccurredOn: req.OccurredOn, DurationSeconds: req.DurationSeconds, SetsCount: req.SetsCount, RepsCount: req.RepsCount, GazeCount: req.GazeCount, CardId: req.CardId, IsCorrect: req.IsCorrect, ReactionMs: req.ReactionMs, Payload: req.Payload, Remark: req.Remark}
}
