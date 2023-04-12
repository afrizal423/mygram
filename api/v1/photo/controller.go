package photo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/afrizal423/mygram/app/business/photo"
	"github.com/afrizal423/mygram/app/models"
	"github.com/afrizal423/mygram/pkg/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Controller struct {
	service photo.IPhotoService
}

func NewPhotoController(service photo.IPhotoService) *Controller {
	return &Controller{
		service,
	}
}

// Get all photo by user id
// User can access to show all photo by user id
// Get all photo godoc
// @Summary Get all photo user
// @Description Get all photo user
// @Tags Photo
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Success 200 {array} models.Photo
// @Failure 500 {object} errors.RestErr
// @Router /photo [get]
func (handler *Controller) GetAllPhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))
	res, err := handler.service.GetAllByUserId(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.NewInternalServerError(
			fmt.Sprintf("Error retrieving all photo user: %s", err.Error())))
		return
	}
	c.JSON(http.StatusOK, res)
}

// Get photo by user id
// User can access to show photo by user id
// Get photo godoc
// @Summary Get photo user
// @Description Get photo user
// @Tags Photo
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param photoId path int true "Photo ID"
// @Success 200 {object} models.Photo
// @Failure 400 {object} errors.RestErr
// @Failure 404 {object} errors.RestErr
// @Router /photo/{photoId} [get]
func (handler *Controller) GetPhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))

	// Get param photo_id
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil || uint(photoId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewBadRequestError("invalid photo ID"))

		return
	}

	res, err := handler.service.GetByUserId(userID, uint(photoId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.NewNotFoundError(fmt.Sprintf("Photo Data with id: %d not found\n", photoId)))
		return
	}
	c.JSON(http.StatusOK, res)
}

// Create photo by user id
// User can create photo by user id
// Create photo godoc
// @Summary Create photo user
// @Description Create photo user
// @Tags Photo
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param requestCreate body models.RequestPhoto true "Create Photo user"
// @Success 201 {object} models.Photo
// @Failure 400 {object} errors.RestErr
// @Router /photo [post]
func (handler *Controller) CreatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))
	Photo := models.Photo{}

	if err := c.ShouldBindJSON(&Photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	}

	if res, err := handler.service.Create(Photo, userID); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	} else {
		c.JSON(http.StatusCreated, res)
		return
	}

}

// Update photo by user id
// User can update photo by user id
// Update photo godoc
// @Summary Update photo user
// @Description Update photo user
// @Tags Photo
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param photoId path int true "Photo ID"
// @Param requestUpdate body models.RequestPhoto true "Update Photo user"
// @Success 200 {object} models.Photo
// @Failure 400 {object} errors.RestErr
// @Router /photo/{photoId} [put]
func (handler *Controller) UpdatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))
	Photo := models.Photo{}

	// Get param photo_id
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil || uint(photoId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewBadRequestError("invalid photo ID"))

		return
	}
	if err := c.ShouldBindJSON(&Photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	}

	if res, err := handler.service.Update(Photo, uint(photoId), userID); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	} else {
		c.JSON(http.StatusCreated, res)
		return
	}
}

// Delete photo by user id
// User can delete photo by user id
// Delete photo godoc
// @Summary Delete photo user
// @Description Delete photo user
// @Tags Photo
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param photoId path int true "Photo ID"
// @Router /photo/{photoId} [delete]
func (handler *Controller) DeletePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))

	// Get Param PhotoId
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil || uint(photoId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewBadRequestError("invalid photo ID"))
		return
	}

	if _, err := handler.service.Delete(uint(photoId), userID); err != nil {
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerError("Failed to delete photo"))
		return
	} else {
		c.JSON(http.StatusOK, ResponseDeleted{
			Message: "Data Photo deleted successfully",
		})
	}
}
