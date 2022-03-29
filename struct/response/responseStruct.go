package responseStruct

import structs "github.com/shivamsouravjha/Micro-Game/struct"

type GetCreatorDetailsResponse struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Data    *structs.UserDetails `json:"userData"`
}

type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
