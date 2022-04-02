package responseStruct

type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type GetContentResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    map[string]string `json:"userData"`
}
