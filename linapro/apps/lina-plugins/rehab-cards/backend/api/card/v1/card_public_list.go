// This file declares the public picture-card list endpoint DTOs. The patient H5
// uses this endpoint to fetch enabled picture cards without requiring any user
// authentication; image rendering is delegated to upstream upload.wikimedia.org
// URLs returned in the response.

package v1

import "github.com/gogf/gf/v2/frame/g"

// PublicListReq defines the request for listing enabled picture cards publicly.
type PublicListReq struct {
	g.Meta     `path:"/rehab/card/public" method:"get" tags:"Rehab Cards" summary:"List enabled picture cards for the patient H5" dc:"Public, unauthenticated endpoint that returns enabled picture cards used by the patient H5 picture-card game. Image URLs in the response point to upstream image hosts; the H5 fetches them directly." `
	CategoryId int64  `json:"categoryId" dc:"Filter by category ID, omitted means all categories" eg:"1"`
	Difficulty int    `json:"difficulty" dc:"Filter by difficulty: 1=easy, 2=normal, 3=hard, omitted means all difficulties" eg:"1"`
	Limit      int    `json:"limit" d:"100" v:"min:1|max:200" dc:"Maximum number of cards to return" eg:"100"`
	Keyword    string `json:"keyword" dc:"Optional fuzzy filter by card title, omitted means no filter" eg:"apple"`
}

// PublicListRes defines the response for the public picture-card list endpoint.
type PublicListRes struct {
	List  []*CardEntity `json:"list" dc:"Public picture-card list" eg:"[]"`
	Total int           `json:"total" dc:"Total number of matched picture cards" eg:"100"`
}
