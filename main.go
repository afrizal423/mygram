package main

import (
	photoController "github.com/afrizal423/mygram/api/v1/photo"
	userController "github.com/afrizal423/mygram/api/v1/user"
	photoService "github.com/afrizal423/mygram/app/business/photo"
	userService "github.com/afrizal423/mygram/app/business/user"
	photoRepository "github.com/afrizal423/mygram/app/repository/photo"
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

	r := router.Route(
		userHandler,
		photoHandler,
	)
	r.Run(":8080")
}
