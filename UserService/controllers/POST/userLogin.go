package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/Micro-Game/UserService/helpers/db"
	requestStruct "github.com/shivamsouravjha/Micro-Game/UserService/struct/request"
	responseStruct "github.com/shivamsouravjha/Micro-Game/UserService/struct/response"
	"github.com/shivamsouravjha/Micro-Game/UserService/utils"
)

func UserLogin(c *gin.Context) {
	userPassword := requestStruct.UserLogin{}
	if err := c.ShouldBind(&userPassword); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	resp := responseStruct.SuccessResponse{}
	token, err := db.UserLoginDAO(c.Request.Context(), &userPassword)
	if err != nil {
		resp.Status = "Failed"
		resp.Message = "Login failed,test penname and password"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status = "Success"
	resp.Message = "user loggedin successfully"
	resp.Token = token
	c.JSON(http.StatusOK, resp)
}
