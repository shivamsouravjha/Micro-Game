package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/Micro-Game/constants"
	"github.com/shivamsouravjha/Micro-Game/helpers/db"
	requestStruct "github.com/shivamsouravjha/Micro-Game/struct/request"
	responseStruct "github.com/shivamsouravjha/Micro-Game/struct/response"
	"github.com/shivamsouravjha/Micro-Game/utils"
)

func CreateUser(c *gin.Context) {
	createUserStruct := requestStruct.CreateUserDetails{}
	if err := c.ShouldBind(&createUserStruct); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	resp := responseStruct.SuccessResponse{}
	err := db.CreateUser(c.Request.Context(), &createUserStruct)
	if err != "" {
		resp.Status = constants.API_FAILED_STATUS
		resp.Message = err
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status = "Success"
	resp.Message = "User Created successfully"
	c.JSON(http.StatusOK, resp)
}
