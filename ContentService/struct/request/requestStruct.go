package requestStruct

import structs "github.com/shivamsouravjha/Micro-Game/ContentService/struct"

type GetContent struct {
	UserID   string `json:"userId" binding:"required"`
	SeriesID string `json:"seriesId" binding:"required"`
}

type ContentUpload struct {
	Data []structs.ContentDetails `json:"content"`
}
