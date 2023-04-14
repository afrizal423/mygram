package router

import (
	"github.com/afrizal423/mygram/api/v1/comment"
	"github.com/afrizal423/mygram/api/v1/photo"
	"github.com/afrizal423/mygram/api/v1/socialmedia"
	"github.com/afrizal423/mygram/api/v1/user"
	_ "github.com/afrizal423/mygram/docs"
	"github.com/afrizal423/mygram/pkg/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title          	Swagger MyGram API
// @version        	1.0
// @description    	This is a sample service for managing MyGram server.
// @contact.name   	Afrizalmy
// @contact.email	afrizalmy1@gmail.com
// @host      		localhost:8080
// @license.name Apache 2.0
// @license.url http://www.apache.org/license/LICENSE-2.0.html
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description description: Enter the token with the `Bearer: ` prefix, e.g. "Bearer abcde12345".
// @externalDocs.description  OpenAPI
func Route(userHandler *user.Controller,
	photoHandler *photo.Controller,
	sosmedHandler *socialmedia.Controller,
	commentHandler *comment.Controller) *gin.Engine {

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

	sosmedRouter := r.Group("/socialmedia")
	{
		sosmedRouter.Use(middlewares.Authentication())
		sosmedRouter.POST("/", sosmedHandler.CreateSosmed)
		sosmedRouter.GET("/", middlewares.SocialMediaAuthorizations(), sosmedHandler.GetAllSosmed)
		sosmedRouter.GET("/:socialMediaId", middlewares.SingleDataSocialMediaAuthorizations(), sosmedHandler.GetSosmed)
		sosmedRouter.PUT("/:socialMediaId", middlewares.SingleDataSocialMediaAuthorizations(), sosmedHandler.UpdateSosmed)
		sosmedRouter.DELETE("/:socialMediaId", middlewares.SingleDataSocialMediaAuthorizations(), sosmedHandler.DeleteSosmed)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", commentHandler.CreateComment)
		commentRouter.GET("/", middlewares.CommentAuthorizations(), commentHandler.GetAllComment)
		commentRouter.GET("/:commentId", middlewares.CommentAuthorizations(), commentHandler.GetComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorizations(), commentHandler.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorizations(), commentHandler.DeleteComment)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
