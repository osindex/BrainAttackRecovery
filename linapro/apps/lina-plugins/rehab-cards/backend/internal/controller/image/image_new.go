// This file wires the image proxy controller dependencies.

package image

import (
	imageapi "lina-plugin-rehab-cards/backend/api/image"
	imagesvc "lina-plugin-rehab-cards/backend/internal/service/image"
)

// ControllerV1 is the Wikimedia image proxy controller.
type ControllerV1 struct{ imageSvc imagesvc.Service }

// NewV1 creates a Wikimedia image proxy controller.
func NewV1() imageapi.IImageV1 { return &ControllerV1{imageSvc: imagesvc.New()} }
