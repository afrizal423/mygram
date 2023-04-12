package router

import (
	"github.com/afrizal423/mygram/api/v1/photo"
	"github.com/afrizal423/mygram/api/v1/user"
	_ "github.com/afrizal423/mygram/docs"
	"github.com/afrizal423/mygram/pkg/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title          	Swagger MyGram API
// @version        	1.0
// @description    	This is a sample server MyGram server.
// @contact.name   	ansharw
// @host      		localhost:8080
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description description: Enter the token with the `Bearer: ` prefix, e.g. "Bearer abcde12345".
// @externalDocs.description  OpenAPI
func Route(userHandler *user.Controller, photoHandler *photo.Controller) *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userHandler.Register)
		userRouter.POST("/login", userHandler.Login)
	}
	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", photoHandler.CreatePhoto)
		photoRouter.GET("/", middlewares.PhotoAuthorizations(), photoHandler.GetAllPhoto)
		photoRouter.GET("/:photoId", middlewares.SingleDataPhotoAuthorizations(), photoHandler.GetPhoto)
		photoRouter.PUT("/:photoId", middlewares.SingleDataPhotoAuthorizations(), photoHandler.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.SingleDataPhotoAuthorizations(), photoHandler.DeletePhoto)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
