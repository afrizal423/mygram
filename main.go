package main

import (
	commentController "github.com/afrizal423/mygram/api/v1/comment"
	photoController "github.com/afrizal423/mygram/api/v1/photo"
	sosmedController "github.com/afrizal423/mygram/api/v1/socialmedia"
	userController "github.com/afrizal423/mygram/api/v1/user"
	commentService "github.com/afrizal423/mygram/app/business/comment"
	photoService "github.com/afrizal423/mygram/app/business/photo"
	sosmedService "github.com/afrizal423/mygram/app/business/socialmedia"
	userService "github.com/afrizal423/mygram/app/business/user"
	commentRepository "github.com/afrizal423/mygram/app/repository/comment"
	photoRepository "github.com/afrizal423/mygram/app/repository/photo"
	sosmedRepository "github.com/afrizal423/mygram/app/repository/socialmedia"
	userRepository "github.com/afrizal423/mygram/app/repository/user"

	"github.com/afrizal423/mygram/configs"
	"github.com/afrizal423/mygram/database"
	"github.com/afrizal423/mygram/router"
)

func main() {
	conn := configs.GormPostgresConn()
	// migrate db
	database.DbMigrate(conn)

	userHandler := userController.NewUserController(userService.NewUserService(
		userRepository.NewuserRepository(conn)))
	photoHandler := photoController.NewPhotoController(photoService.NewPhotoService(
		photoRepository.NewPhotoRepository(conn)))
	sosmedHandler := sosmedController.NewSocialMediaController(sosmedService.NewSocialMediaService(
		sosmedRepository.NewSocialMediaRepository(conn)))

	commentHandler := commentController.NewCommentController(commentService.NewCommentService(
		commentRepository.NewCommentRepository(conn)))

	r := router.Route(
		userHandler,
		photoHandler,
		sosmedHandler,
		commentHandler,
	)
	var PORT = configs.Config("PORT")
	r.Run(":" + PORT)
}
