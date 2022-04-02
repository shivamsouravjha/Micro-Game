package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/Micro-Game/ContentService/constants"
	"github.com/shivamsouravjha/Micro-Game/ContentService/helpers/db"
	requestStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/request"
	responseStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/response"
)

func GetSeries(c *gin.Context) {
	seriesId := c.Params.ByName("seriesId")
	getSeriesRequest := requestStruct.GetSeries{
		SeriesID: seriesId,
	}
	resp := responseStruct.GetSeriesDetails{}
	UserDetails, err := db.GetSeriesDAO(c.Request.Context(), &getSeriesRequest)
	if err != nil {
		resp.Status = constants.API_FAILED_STATUS
		resp.Message = "Fetching Content Failed"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status = "Success"
	resp.Message = "Series fetched successfully"
	resp.Data = UserDetails
	c.JSON(http.StatusOK, resp)
}
