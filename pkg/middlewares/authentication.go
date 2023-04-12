package middlewares

import (
	"fmt"
	"net/http"

	tokenjwt "github.com/afrizal423/mygram/pkg/utils/tokenJWT"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		dt, err := tokenjwt.TokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", dt)
		fmt.Println(dt)
		c.Next()
	}
}
