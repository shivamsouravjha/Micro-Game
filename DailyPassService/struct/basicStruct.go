package structs

type ChapterDetails struct {
	UnlockedContent map[string]interface{} `json:"unlockedSeries"`
}

type NewUserContent struct {
	UserId           string                 `json:"userId" binding:"required"`
	UnlockedChapters map[string]interface{} `json:"unlockedChapters" binding:"required"`
}

type UserContent struct {
	Chapter      string `json:"chapter" binding:"required"`
	ChapterIndex string `json:"chapterIndex" binding:"required"`
}

type NewChapterContent struct {
	UserId           string `json:"userId" binding:"required"`
	UnlockedChapters string `json:"unlockedChapters" binding:"required"`
}
