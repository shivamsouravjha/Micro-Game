package routes

import (
	"github.com/gin-gonic/gin"
	GET "github.com/shivamsouravjha/Micro-Game/DailyPassService/controllers/GET"
	POST "github.com/shivamsouravjha/Micro-Game/DailyPassService/controllers/POST"
)

func v0Routes(route *gin.RouterGroup) {
	v1Routes := route.Group("/v0")
	{
		v1Routes.GET("/dailypass/getUnlockedContent/:userId/", GET.GetUnlockedContent)
		v1Routes.POST("/dailypass/unlockContent", POST.UnlockContent)
	}
}
