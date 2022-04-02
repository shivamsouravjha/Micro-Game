package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/Micro-Game/ContentService/helpers/db"
	requestStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/request"
	responseStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/response"
	"github.com/shivamsouravjha/Micro-Game/ContentService/utils"
)

func UploadSeries(c *gin.Context) {
	createSeriesStruct := requestStruct.SeriesUpload{}
	if err := c.ShouldBind(&createSeriesStruct); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	resp := responseStruct.SuccessResponse{}
	db.UploadSeries(c.Request.Context(), &createSeriesStruct)
	resp.Status = "Success"
	resp.Message = "Content Uploaded successfully"
	c.JSON(http.StatusOK, resp)
}
