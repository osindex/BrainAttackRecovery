// This file wires the category controller dependencies.

package category

import (
	categoryapi "lina-plugin-rehab-cards/backend/api/category"
	v1 "lina-plugin-rehab-cards/backend/api/category/v1"
	categorysvc "lina-plugin-rehab-cards/backend/internal/service/category"
)

// ControllerV1 is the picture-card category controller.
type ControllerV1 struct{ categorySvc categorysvc.Service }

// NewV1 creates a picture-card category controller.
func NewV1() categoryapi.ICategoryV1 { return &ControllerV1{categorySvc: categorysvc.New()} }

// toAPIEntity converts a service entity to an API projection.
func toAPIEntity(e *categorysvc.Entity) *v1.CategoryEntity {
	if e == nil {
		return nil
	}
	return &v1.CategoryEntity{Id: e.Id, Name: e.Name, Code: e.Code, Sort: e.Sort, Status: e.Status, Remark: e.Remark, CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt}
}
