package requestStruct

type GetUserDetailsRequest struct {
	PenName string `json:"penName" binding:"required"`
}

type GetUnlockedContent struct {
	UserId   int `json:"userId" binding:"required"`
	SeriesId int `json:"seriesId" binding:"required"`
}

type CreateUserDetails struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"`
	PenName   string `json:"penName" binding:"required"`
	UserEmail string `json:"userEmail" binding:"required"`
	Bio       string `json:"bio"`
	Number    string `json:"number"`
}

type UnlockContent struct {
	UserId    int   `json:"userId"`
	ContentID []int `json:"contentId"`
	SeriesID  int   `json:"seriesId"`
}
