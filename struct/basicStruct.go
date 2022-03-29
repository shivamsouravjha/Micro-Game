package structs

type UserDetails struct {
	UserId    int    `form:"userId" binding:"required"`
	FirstName string `form:"firstName" binding:"required"`
	LastName  string `form:"lastName"`
	PenName   string `form:"penName" binding:"required"`
	UserEmail string `form:"userEmail" binding:"required"`
	Bio       string `form:"bio"`
	Number    string `form:"number"`
}
