package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/Micro-Game/constants"
	"github.com/shivamsouravjha/Micro-Game/helpers/db"
	requestStruct "github.com/shivamsouravjha/Micro-Game/struct/request"
	responseStruct "github.com/shivamsouravjha/Micro-Game/struct/response"
	"github.com/shivamsouravjha/Micro-Game/utils"
)

func GetUserDetails(c *gin.Context) {
	getCreatorRequest := requestStruct.GetUserDetailsRequest{}
	if err := c.ShouldBind(&getCreatorRequest); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	resp := responseStruct.GetCreatorDetailsResponse{}

	UserDetails, err := db.GetCreator(c.Request.Context(), &getCreatorRequest)
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
