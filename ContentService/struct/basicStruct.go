package structs

type ContentDetails struct {
	SeriesID string `form:"seriesId" binding:"required"`
	Title    string `form:"title" binding:"required"`
	Story    string `form:"story" binding:"required"`
}

type SeriesDetails struct {
	SeriesID string `form:"seriesId" binding:"required"`
	Author   string `form:"author" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Chapters string `form:"chapters" binding:"required"`
}
