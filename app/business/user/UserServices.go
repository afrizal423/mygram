package user

import (
	"errors"
	"fmt"

	"github.com/afrizal423/mygram/app/models"
	"github.com/afrizal423/mygram/pkg/utils/hash"
	tokenjwt "github.com/afrizal423/mygram/pkg/utils/tokenJWT"
)

type UserService struct {
	Repository IUserRepository
	Hashing    hash.Hashing
}

func NewUserService(repository IUserRepository) *UserService {
	return &UserService{
		repository,
		hash.Hashing{},
	}
}

func (u *UserService) Register(data models.User) (models.User, error) {
	var dt models.User
	// hash password
	// data.Password = u.hashing.HashPassword(data.Password)
	// repositoty to insert data
	data, err := u.Repository.Register(data)
	if err != nil {
		return dt, err
	}
	return data, nil
}

func (u *UserService) Login(data models.User) (string, error) {
	dt, err := u.Repository.FindByEmail(data.Email)
	fmt.Println(data)
	if err != nil {
		return "", errors.New("tidak ada data")
	}
	// fmt.Println(data.Password, dt.Password)
	// return "", nil
	match, err := u.Hashing.VerifikasiPassword(data.Password, dt.Password)
	if err != nil {
		return "", err
	}

	if match {
		// bener
		token, _ := tokenjwt.GenerateToken(dt.ID, dt.Email, dt.Username)
		return token, nil
	} else {
		// salah
		return "", errors.New("password salah")
	}
}
