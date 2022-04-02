package routes

import (
	"github.com/gin-gonic/gin"
	GET "github.com/shivamsouravjha/Micro-Game/UserService/controllers/GET"
	POST "github.com/shivamsouravjha/Micro-Game/UserService/controllers/POST"
)

func v0Routes(route *gin.RouterGroup) {
	v1Routes := route.Group("/v0")
	{
		v1Routes.GET("/getUser/:penName", GET.GetUserDetails)
		v1Routes.POST("/createUser", POST.CreateUser)
	}
}
