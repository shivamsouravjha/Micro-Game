package routes

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.New()
	v1 := router.Group("/api")
	v0Routes(v1)

	return router
}
