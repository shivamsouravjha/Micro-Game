package routes

import (
	"github.com/gin-gonic/gin"
	GET "github.com/shivamsouravjha/Micro-Game/ContentService/controllers/GET"
	POST "github.com/shivamsouravjha/Micro-Game/ContentService/controllers/POST"
)

func v0Routes(route *gin.RouterGroup) {
	v1Routes := route.Group("/v0/content")
	{
		v1Routes.GET("/fetchContent", GET.GetContent)
		v1Routes.POST("/uploadContent", POST.UploadContent)
		v1Routes.GET("/getSeries", GET.GetSeries)
		v1Routes.POST("/uploadSeries", POST.UploadSeries)
	}
}
