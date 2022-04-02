package requestStruct

import structs "github.com/shivamsouravjha/Micro-Game/ContentService/struct"

type GetContent struct {
	UserID   string `json:"userId" binding:"required"`
	SeriesID string `json:"seriesId" binding:"required"`
}

type GetSeries struct {
	SeriesID string `json:"seriesId" binding:"required"`
}
type ContentUpload struct {
	Data []structs.ContentDetails `json:"content"`
}
type SeriesUpload struct {
	Author   string   `form:"author" binding:"required"`
	Name     string   `form:"name" binding:"required"`
	Chapters []string `form:"chapters" binding:"required"`
}
