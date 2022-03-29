package requestStruct

type GetUserDetailsRequest struct {
	PenName string `json:"penName" binding:"required"`
}
