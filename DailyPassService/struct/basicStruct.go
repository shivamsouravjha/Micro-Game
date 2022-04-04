package structs

type ChapterDetails struct {
	UnlockedContent map[string]interface{} `json:"unlockedSeries"`
}

type NewUserContent struct {
	UserId           string                 `json:"userId" binding:"required"`
	UnlockedChapters map[string]interface{} `json:"unlockedChapters" binding:"required"`
}
