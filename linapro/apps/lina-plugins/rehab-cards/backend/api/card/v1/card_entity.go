// This file declares card response entities for the rehab-cards plugin.

package v1

import "github.com/gogf/gf/v2/os/gtime"

// CardEntity describes one picture card returned by the API.
type CardEntity struct {
	Id           int64       `json:"id" dc:"Card ID" eg:"1"`
	CategoryId   int64       `json:"categoryId" dc:"Category ID" eg:"1"`
	CategoryName string      `json:"categoryName" dc:"Category display name" eg:"Animals"`
	Title        string      `json:"title" dc:"Card title" eg:"Apple"`
	Label        string      `json:"label" dc:"Expected answer label" eg:"apple"`
	ImageUrl     string      `json:"imageUrl" dc:"Image URL served to the H5 app" eg:"/api/v1/uploads/2026/05/apple.png"`
	ImageFileId  int64       `json:"imageFileId" dc:"Uploaded file ID, 0 means the card uses an external or placeholder URL" eg:"1"`
	Source       string      `json:"source" dc:"Image source: manual, mock, pixabay, wikimedia, or another provider" eg:"manual"`
	SourceUrl    string      `json:"sourceUrl" dc:"Original source URL used for attribution or troubleshooting" eg:"https://example.com/apple"`
	License      string      `json:"license" dc:"Image license label" eg:"CC0"`
	Difficulty   int         `json:"difficulty" dc:"Difficulty: 1=easy, 2=normal, 3=hard" eg:"1"`
	Status       int         `json:"status" dc:"Status: 1=enabled, 0=disabled" eg:"1"`
	Sort         int         `json:"sort" dc:"Sort weight, smaller values are displayed earlier" eg:"10"`
	Remark       string      `json:"remark" dc:"Optional remark" eg:"High-frequency noun"`
	CreatedAt    *gtime.Time `json:"createdAt" dc:"Creation time" eg:"2026-05-09 10:00:00"`
	UpdatedAt    *gtime.Time `json:"updatedAt" dc:"Last update time" eg:"2026-05-09 10:00:00"`
}
