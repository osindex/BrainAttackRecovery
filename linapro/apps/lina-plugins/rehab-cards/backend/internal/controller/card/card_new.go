// This file wires the card controller dependencies.

package card

import (
	cardapi "lina-plugin-rehab-cards/backend/api/card"
	v1 "lina-plugin-rehab-cards/backend/api/card/v1"
	cardsvc "lina-plugin-rehab-cards/backend/internal/service/card"
)

// ControllerV1 is the picture-card controller.
type ControllerV1 struct{ cardSvc cardsvc.Service }

// NewV1 creates a picture-card controller.
func NewV1() cardapi.ICardV1 { return &ControllerV1{cardSvc: cardsvc.New()} }

// toAPIEntity converts a service entity to an API projection.
func toAPIEntity(e *cardsvc.Entity) *v1.CardEntity {
	if e == nil {
		return nil
	}
	return &v1.CardEntity{Id: e.Id, CategoryId: e.CategoryId, CategoryName: e.CategoryName, Title: e.Title, Label: e.Label, ImageUrl: e.ImageUrl, ImageFileId: e.ImageFileId, Source: e.Source, SourceUrl: e.SourceUrl, License: e.License, Difficulty: e.Difficulty, Status: e.Status, Sort: e.Sort, Remark: e.Remark, CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt}
}

// saveInputFromCreate converts a create request to service input.
func saveInputFromCreate(req *v1.CreateReq) cardsvc.SaveInput {
	return cardsvc.SaveInput{CategoryId: req.CategoryId, Title: req.Title, Label: req.Label, ImageUrl: req.ImageUrl, ImageFileId: req.ImageFileId, Source: req.Source, SourceUrl: req.SourceUrl, License: req.License, Difficulty: req.Difficulty, Status: req.Status, Sort: req.Sort, Remark: req.Remark}
}

// saveInputFromUpdate converts an update request to service input.
func saveInputFromUpdate(req *v1.UpdateReq) cardsvc.SaveInput {
	return cardsvc.SaveInput{CategoryId: req.CategoryId, Title: req.Title, Label: req.Label, ImageUrl: req.ImageUrl, ImageFileId: req.ImageFileId, Source: req.Source, SourceUrl: req.SourceUrl, License: req.License, Difficulty: req.Difficulty, Status: req.Status, Sort: req.Sort, Remark: req.Remark}
}
