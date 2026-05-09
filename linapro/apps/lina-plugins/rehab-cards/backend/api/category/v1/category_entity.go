// This file declares category response entities for the rehab-cards plugin.

package v1

import "github.com/gogf/gf/v2/os/gtime"

// CategoryEntity describes one picture-card category returned by the API.
type CategoryEntity struct {
	Id        int64       `json:"id" dc:"Category ID" eg:"1"`
	Name      string      `json:"name" dc:"Category display name" eg:"Animals"`
	Code      string      `json:"code" dc:"Stable category code" eg:"animal"`
	Sort      int         `json:"sort" dc:"Sort weight, smaller values are displayed earlier" eg:"10"`
	Status    int         `json:"status" dc:"Status: 1=enabled, 0=disabled" eg:"1"`
	Remark    string      `json:"remark" dc:"Optional remark" eg:"Common object naming category"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"Creation time" eg:"2026-05-09 10:00:00"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"Last update time" eg:"2026-05-09 10:00:00"`
}
