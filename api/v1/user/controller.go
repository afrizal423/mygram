package user

import (
	"fmt"
	"net/http"

	"github.com/afrizal423/mygram/app/business/user"
	"github.com/afrizal423/mygram/app/models"
	"github.com/afrizal423/mygram/pkg/utils/errors"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service user.IUserService
}

func NewUserController(service user.IUserService) *Controller {
	return &Controller{
		service,
	}
}

// Register user
// User can register account
// Register godoc
// @Summary Register user
// @Description Register user
// @Tags User
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Param requestRegister body models.UserRegister true "Register user"
// @Failure 400 {object} errors.RestErr
// @Router /users/register [post]
func (c *Controller) Register(ctx *gin.Context) {
	// alokasikan memori
	// https://dev.to/bhanu011/how-is-new-different-in-go-41lk
	// https://go.dev/doc/effective_go
	user := new(models.User)
	// pastikan data harus berupa json
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}

	// proses ke bagian services
	data, err := c.service.Register(*user)
	if err != nil {
		restErr := errors.NewBadRequestError(fmt.Sprintf("%v", err.Error()))
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}
	// tinggal di return aja
	ctx.JSON(http.StatusCreated, gin.H{
		"id":       data.ID,
		"username": data.Username,
		"email":    data.Email,
		"age":      data.Age})
}

// Login
// User can login account
// Login godoc
// @Summary Login user
// @Description Login user
// @Tags User
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Param requestLogin body models.UserLogin true "login user"
// @Failure 400 {object} errors.RestErr
// @Failure 401 {object} errors.RestErr
// @Router /users/login [post]
func (c *Controller) Login(ctx *gin.Context) {
	// user := new(models.User)
	var user models.User
	// pastikan data harus berupa json
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}
	token, err := c.service.Login(user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errors.NewUnauthorizedError(err.Error()))
		return
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"token": token,
		})
	}
}
