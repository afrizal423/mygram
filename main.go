package main

import (
	userController "github.com/afrizal423/mygram/api/v1/user"
	userService "github.com/afrizal423/mygram/app/business/user"
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

	r := router.Route(
		userHandler)
	r.Run(":8080")
}
