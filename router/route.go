package router

import (
	"github.com/afrizal423/mygram/api/v1/user"
	"github.com/gin-gonic/gin"
)

func Route(userHandler *user.Controller) *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userHandler.Register)
		userRouter.POST("/login", userHandler.Login)
	}
	return r
}
