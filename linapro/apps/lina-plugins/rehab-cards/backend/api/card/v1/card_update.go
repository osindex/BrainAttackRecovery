// This file declares the update-card endpoint DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// UpdateReq defines the request for updating a picture card.
type UpdateReq struct {
	g.Meta      `path:"/rehab/card/{id}" method:"put" tags:"Rehab Cards" summary:"Update picture card" dc:"Update one picture card by ID." permission:"rehab:card:edit"`
	Id          int64  `json:"id" v:"required|min:1" dc:"Card ID" eg:"1"`
	CategoryId  int64  `json:"categoryId" v:"required|min:1" dc:"Category ID" eg:"1"`
	Title       string `json:"title" v:"required|max-length:100" dc:"Card title" eg:"Apple"`
	Label       string `json:"label" v:"required|max-length:100" dc:"Expected answer label" eg:"apple"`
	ImageUrl    string `json:"imageUrl" dc:"Image URL served to the H5 app" eg:"/api/v1/uploads/2026/05/apple.png"`
	ImageFileId int64  `json:"imageFileId" dc:"Uploaded file ID, 0 means no uploaded file is attached" eg:"1"`
	Source      string `json:"source" dc:"Image source: manual, mock, pixabay, wikimedia, or another provider" eg:"manual"`
	SourceUrl   string `json:"sourceUrl" dc:"Original source URL used for attribution or troubleshooting" eg:"https://example.com/apple"`
	License     string `json:"license" dc:"Image license label" eg:"CC0"`
	Difficulty  int    `json:"difficulty" dc:"Difficulty: 1=easy, 2=normal, 3=hard" eg:"1"`
	Status      int    `json:"status" dc:"Status: 1=enabled, 0=disabled" eg:"1"`
	Sort        int    `json:"sort" dc:"Sort weight, smaller values are displayed earlier" eg:"10"`
	Remark      string `json:"remark" dc:"Optional remark" eg:"High-frequency noun"`
}

// UpdateRes defines the response for updating a picture card.
type UpdateRes struct{}
