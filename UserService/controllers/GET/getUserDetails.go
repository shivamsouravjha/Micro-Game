package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/Micro-Game/UserService/constants"
	"github.com/shivamsouravjha/Micro-Game/UserService/helpers/db"
	responseStruct "github.com/shivamsouravjha/Micro-Game/UserService/struct/response"
)

func GetUserDetails(c *gin.Context) {
	penName := c.Params.ByName("penName")
	resp := responseStruct.GetCreatorDetailsResponse{}
	UserDetails, err := db.GetCreatorDAO(c.Request.Context(), penName)
	if err != nil {
		resp.Status = constants.API_FAILED_STATUS
		resp.Message = "Fetching Details Failed"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status = "Success"
	resp.Message = "Details fetched successfully"
	resp.Data = UserDetails
	c.JSON(http.StatusOK, resp)
}
