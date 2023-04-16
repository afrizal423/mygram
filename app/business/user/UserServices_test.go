package user_test

import (
	"fmt"
	"testing"

	"github.com/afrizal423/mygram/app/business/user"
	"github.com/afrizal423/mygram/app/business/user/mocks"
	"github.com/afrizal423/mygram/app/models"
	"github.com/afrizal423/mygram/pkg/utils/hash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepoInterfaceMock = &mocks.IUserRepository{Mock: mock.Mock{}}
var userUseCaseInterface = user.UserService{Repository: userRepoInterfaceMock, Hashing: hash.Hashing{}}

func TestRegister(t *testing.T) {
	//data mock hasil login
	userDataDummyLogin := models.User{
		Username: "LeeMark",
		Email:    "leemark@gmail.com",
		Password: "123456",
		Age:      9,
	}
	userRepoInterfaceMock.On("Register", userDataDummyLogin).Return(userDataDummyLogin, nil).Once()
	requestLoginUser := models.User{
		Username: "LeeMark",
		Email:    "leemark@gmail.com",
		Password: "123456",
		Age:      9,
	}
	user, err := userUseCaseInterface.Register(requestLoginUser)
	fmt.Println(user)
	assert.Equal(t, nil, err)
	assert.Equal(t, userDataDummyLogin, user)
}
