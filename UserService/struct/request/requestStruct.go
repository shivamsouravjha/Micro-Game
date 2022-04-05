package requestStruct

type GetUserDetailsRequest struct {
	PenName string `json:"penName" binding:"required"`
}

type UserLogin struct {
	Password string `json:"password" binding:"required"`
	PenName  string `json:"penName" binding:"required"`
}
type GetUnlockedContent struct {
	UserId   string `json:"userId" binding:"required"`
	SeriesId string `json:"seriesId" binding:"required"`
}

type CreateUserDetails struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"`
	PenName   string `json:"penName" binding:"required"`
	UserEmail string `json:"userEmail" binding:"required"`
	Bio       string `json:"bio"`
	Number    string `json:"number"`
	Password  string `json:"password"`
}

type UnlockContent struct {
	UserId    string   `json:"userId"`
	ContentID []string `json:"contentId"`
	SeriesID  string   `json:"seriesId"`
}
