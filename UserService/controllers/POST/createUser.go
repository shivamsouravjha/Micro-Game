package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/Micro-Game/UserService/constants"
	"github.com/shivamsouravjha/Micro-Game/UserService/helpers/db"
	requestStruct "github.com/shivamsouravjha/Micro-Game/UserService/struct/request"
	responseStruct "github.com/shivamsouravjha/Micro-Game/UserService/struct/response"
	"github.com/shivamsouravjha/Micro-Game/UserService/utils"
)

func CreateUser(c *gin.Context) {
	createUserStruct := requestStruct.CreateUserDetails{}
	if err := c.ShouldBind(&createUserStruct); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	resp := responseStruct.SuccessResponse{}
	err := db.CreateUserDAO(c.Request.Context(), &createUserStruct)
	if err != "" {
		resp.Status = constants.API_FAILED_STATUS
		resp.Message = err
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	token, tokenerror := utils.GenerateToken()

	if tokenerror != nil {
		resp.Status = constants.API_FAILED_STATUS
		resp.Message = "User Created,Please login"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status = "Success"
	resp.Message = "User Created successfully"
	resp.Token = token
	c.JSON(http.StatusOK, resp)
}
