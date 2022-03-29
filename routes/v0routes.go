package routes

import (
	"github.com/gin-gonic/gin"
	GET "github.com/shivamsouravjha/Micro-Game/controllers/GET"
	POST "github.com/shivamsouravjha/Micro-Game/controllers/POST"
)

func v0Routes(route *gin.RouterGroup) {
	v1Routes := route.Group("/v0")
	{
		v1Routes.GET("/getCreator", GET.GetUserDetails)
		v1Routes.POST("/getCreatorFeed", POST.CreateUser)
	}
}
