package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/Micro-Game/DailyPassService/constants"
	"github.com/shivamsouravjha/Micro-Game/DailyPassService/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("userToken")
		if token != "" {
			jwtCheck, err := utils.TokenParse(token, "mySigningKey")
			if err != nil {
				c.AbortWithStatusJSON(401, constants.INVALID_TOKEN_RESPONSE)
				return
			} else if jwtCheck.Valid {
				c.Next()
			}
		} else {
			c.AbortWithStatusJSON(401, constants.INVALID_TOKEN_RESPONSE)
		}
	}
}
