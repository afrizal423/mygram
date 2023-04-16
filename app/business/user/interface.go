package user

import "github.com/afrizal423/mygram/app/models"

type IUserService interface {
	Register(data models.User) (models.User, error)
	Login(data models.User) (string, error)
}

//go:generate mockery --name IUserRepository
type IUserRepository interface {
	Register(data models.User) (models.User, error)
	FindByEmail(email string) (*models.User, error)
}
