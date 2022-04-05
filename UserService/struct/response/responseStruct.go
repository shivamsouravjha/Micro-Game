package responseStruct

import structs "github.com/shivamsouravjha/Micro-Game/UserService/struct"

type GetCreatorDetailsResponse struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Data    *structs.UserDetails `json:"userData"`
}

type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

type GetContentResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []interface{} `json:"userData"`
}
