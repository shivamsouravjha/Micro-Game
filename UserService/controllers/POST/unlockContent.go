package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/Micro-Game/UserService/helpers/db"
	requestStruct "github.com/shivamsouravjha/Micro-Game/UserService/struct/request"
	responseStruct "github.com/shivamsouravjha/Micro-Game/UserService/struct/response"
	"github.com/shivamsouravjha/Micro-Game/UserService/utils"
)

func UnlockContent(c *gin.Context) {
	unlockContent := requestStruct.UnlockContent{}
	if err := c.ShouldBind(&unlockContent); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	resp := responseStruct.SuccessResponse{}
	db.UnlockContentDAO(c.Request.Context(), &unlockContent)
	resp.Status = "Success"
	resp.Message = "Content Unlocked successfully"
	c.JSON(http.StatusOK, resp)
}
