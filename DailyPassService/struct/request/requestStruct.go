package requestStruct

type GetUnlockedContent struct {
	UserId   string `json:"userId" binding:"required"`
	SeriesId string `json:"seriesId" binding:"required"`
}

type UnlockContent struct {
	UserId    string   `json:"userId"`
	ContentID []string `json:"contentId"`
	SeriesID  string   `json:"seriesId"`
}
