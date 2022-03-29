package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
	requestStruct "github.com/shivamsouravjha/Micro-Game/struct/request"
	"github.com/shivamsouravjha/Micro-Game/utils"
)

func GetUserDetails(c *gin.Context) {
	getCreatorRequest := requestStruct.GetUserDetailsRequest{}
	if err := c.ShouldBind(&getCreatorRequest); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, getCreatorRequest)
}
