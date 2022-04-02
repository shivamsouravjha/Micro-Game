package structs

type ContentDetails struct {
	SeriesID string `form:"seriesId" binding:"required"`
	Title    string `form:"title" binding:"required"`
	Story    string `form:"story" binding:"required"`
}
