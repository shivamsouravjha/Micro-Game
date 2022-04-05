package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/Micro-Game/ContentService/constants"
	"github.com/shivamsouravjha/Micro-Game/ContentService/helpers/db"
	requestStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/request"
	responseStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/response"
)

func GetContent(c *gin.Context) {
	seriesId := c.Params.ByName("seriesId")
	userId := c.Params.ByName("userId")
	getUserRequest := requestStruct.GetContent{
		SeriesID: seriesId,
		UserID:   userId,
	}
	resp := responseStruct.GetContentDetails{}
	UserDetails, err := db.GetContentDAO(c.Request.Context(), &getUserRequest)
	if err != nil {
		resp.Status = constants.API_FAILED_STATUS
		resp.Message = "Fetching Content Failed"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status = "Success"
	resp.Message = "Content fetched successfully"
	resp.Data = UserDetails
	c.JSON(http.StatusOK, resp)
}
