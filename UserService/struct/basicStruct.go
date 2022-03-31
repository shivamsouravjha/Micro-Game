package structs

type UserDetails struct {
	UserId         int           `json:"userId" binding:"required"`
	FirstName      string        `json:"firstName" binding:"required"`
	LastName       string        `json:"lastName"`
	PenName        string        `json:"penName" binding:"required"`
	UserEmail      string        `json:"userEmail" binding:"required"`
	Bio            string        `json:"bio"`
	Number         string        `json:"number"`
	UnlockedSeries map[int][]int `json:"unlockedSeries"`
}

type ChapterDetails struct {
	ChapterId map[string]string `json:"chapterId"`
}
