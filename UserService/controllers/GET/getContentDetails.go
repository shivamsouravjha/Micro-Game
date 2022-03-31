package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/Micro-Game/UserService/constants"
	"github.com/shivamsouravjha/Micro-Game/UserService/helpers/db"
	requestStruct "github.com/shivamsouravjha/Micro-Game/UserService/struct/request"
	responseStruct "github.com/shivamsouravjha/Micro-Game/UserService/struct/response"
	"github.com/shivamsouravjha/Micro-Game/UserService/utils"
)

func GetUnlockedContent(c *gin.Context) {
	getContentRequest := requestStruct.GetUnlockedContent{}
	if err := c.ShouldBindJSON(&getContentRequest); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	resp := responseStruct.GetCreatorDetailsResponse{}
	UserDetails, err := db.GetContentDAO(c.Request.Context(), &getContentRequest)
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
