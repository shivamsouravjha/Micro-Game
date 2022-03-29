package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
	requestStruct "github.com/shivamsouravjha/Micro-Game/struct/request"
)

func GetUserDetails(c *gin.Context) {
	getCreatorRequest := requestStruct.GetUserDetailsRequest{}

	c.JSON(http.StatusOK, getCreatorRequest)
}
