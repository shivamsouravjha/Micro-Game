package requestStruct

type GetUserDetailsRequest struct {
	PenName string `json:"penName" binding:"required"`
}

type CreateUserDetails struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"`
	PenName   string `json:"penName" binding:"required"`
	UserEmail string `json:"userEmail" binding:"required"`
	Bio       string `json:"bio"`
	Number    string `json:"number"`
}
