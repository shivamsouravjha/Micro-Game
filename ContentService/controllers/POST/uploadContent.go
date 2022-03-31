package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/Micro-Game/ContentService/helpers/db"
	requestStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/request"
	responseStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/response"
	"github.com/shivamsouravjha/Micro-Game/ContentService/utils"
)

func UploadContent(c *gin.Context) {
	createContentStruct := requestStruct.ContentUpload{}
	if err := c.ShouldBind(&createContentStruct); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	resp := responseStruct.SuccessResponse{}
	db.UploadContent(c.Request.Context(), &createContentStruct)
	resp.Status = "Success"
	resp.Message = "Content Uploaded successfully"
	c.JSON(http.StatusOK, resp)
}
