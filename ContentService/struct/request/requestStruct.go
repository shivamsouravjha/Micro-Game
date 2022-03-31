package requestStruct

import structs "github.com/shivamsouravjha/Micro-Game/ContentService/struct"

type GetContent struct {
	UserID   int `json:"userId" binding:"required"`
	SeriesID int `json:"seriesId" binding:"required"`
}

type ContentUpload struct {
	Data []structs.ContentDetails `json:"content"`
}
