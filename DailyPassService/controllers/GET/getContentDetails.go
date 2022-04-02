package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/Micro-Game/DailyPassService/constants"
	"github.com/shivamsouravjha/Micro-Game/DailyPassService/helpers/db"
	requestStruct "github.com/shivamsouravjha/Micro-Game/DailyPassService/struct/request"
	responseStruct "github.com/shivamsouravjha/Micro-Game/DailyPassService/struct/response"
)

func GetUnlockedContent(c *gin.Context) {
	userId := c.Params.ByName("userId")
	seriesId := c.Params.ByName("seriesId")
	getContentRequest := requestStruct.GetUnlockedContent{
		UserId:   userId,
		SeriesId: seriesId,
	}
	resp := responseStruct.GetContentResponse{}
	UserContentList, err := db.GetContentDAO(c.Request.Context(), &getContentRequest)
	userDetails := UserContentList[getContentRequest.SeriesId]
	if err != nil {
		resp.Status = constants.API_FAILED_STATUS
		resp.Message = "Unlocking Content Failed"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status = "Success"
	resp.Message = "Contents unlocked successfully"
	resp.Data = userDetails
	c.JSON(http.StatusOK, resp)
}
