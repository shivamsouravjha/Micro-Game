package structs

type ContentDetails struct {
	SeriesID string `json:"seriesId" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Story    string `json:"story" binding:"required"`
}

type SeriesDetails struct {
	SeriesID string `json:"seriesId" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Chapters string `json:"chapters" binding:"required"`
}

type NewUserDetails struct {
	UserId           string                 `json:"userId" binding:"required"`
	UnlockedChapters map[string]interface{} `json:"unlockedChapters" binding:"required"`
}

type UserContent struct {
	Chapter      string `json:"chapter" binding:"required"`
	ChapterIndex string `json:"chapterIndex" binding:"required"`
}
