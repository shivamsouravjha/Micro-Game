package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/Micro-Game/DailyPassService/helpers/db"
	requestStruct "github.com/shivamsouravjha/Micro-Game/DailyPassService/struct/request"
	responseStruct "github.com/shivamsouravjha/Micro-Game/DailyPassService/struct/response"
	"github.com/shivamsouravjha/Micro-Game/DailyPassService/utils"
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
