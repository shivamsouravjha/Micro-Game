package responseStruct

import structs "github.com/shivamsouravjha/Micro-Game/ContentService/struct"

type GetContentDetails struct {
	Status  string                    `json:"status"`
	Message string                    `json:"message"`
	Data    *[]structs.ContentDetails `json:"metaData"`
}

type GetSeriesDetails struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    *structs.SeriesDetails `json:"metaData"`
}
type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
