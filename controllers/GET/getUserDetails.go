package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserDetails(c *gin.Context) {
	c.JSON(http.StatusOK, "resp")
}
