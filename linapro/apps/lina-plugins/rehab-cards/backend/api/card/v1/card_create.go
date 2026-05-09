// This file declares the create-card endpoint DTOs.

package v1

import "github.com/gogf/gf/v2/frame/g"

// CreateReq defines the request for creating a picture card.
type CreateReq struct {
	g.Meta      `path:"/rehab/card" method:"post" tags:"Rehab Cards" summary:"Create picture card" dc:"Create one picture card used by the H5 object naming mini-game." permission:"rehab:card:add"`
	CategoryId  int64  `json:"categoryId" v:"required|min:1" dc:"Category ID" eg:"1"`
	Title       string `json:"title" v:"required|max-length:100" dc:"Card title" eg:"Apple"`
	Label       string `json:"label" v:"required|max-length:100" dc:"Expected answer label" eg:"apple"`
	ImageUrl    string `json:"imageUrl" dc:"Image URL served to the H5 app" eg:"/api/v1/uploads/2026/05/apple.png"`
	ImageFileId int64  `json:"imageFileId" dc:"Uploaded file ID, 0 means no uploaded file is attached" eg:"1"`
	Source      string `json:"source" d:"manual" dc:"Image source: manual, mock, pixabay, wikimedia, or another provider" eg:"manual"`
	SourceUrl   string `json:"sourceUrl" dc:"Original source URL used for attribution or troubleshooting" eg:"https://example.com/apple"`
	License     string `json:"license" dc:"Image license label" eg:"CC0"`
	Difficulty  int    `json:"difficulty" d:"1" dc:"Difficulty: 1=easy, 2=normal, 3=hard" eg:"1"`
	Status      int    `json:"status" d:"1" dc:"Status: 1=enabled, 0=disabled" eg:"1"`
	Sort        int    `json:"sort" dc:"Sort weight, smaller values are displayed earlier" eg:"10"`
	Remark      string `json:"remark" dc:"Optional remark" eg:"High-frequency noun"`
}

// CreateRes defines the response for creating a picture card.
type CreateRes struct {
	Id int64 `json:"id" dc:"Created card ID" eg:"1"`
}
