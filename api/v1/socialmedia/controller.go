package socialmedia

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/afrizal423/mygram/app/business/socialmedia"
	"github.com/afrizal423/mygram/app/models"
	"github.com/afrizal423/mygram/pkg/utils/errors"
	"github.com/afrizal423/mygram/pkg/utils/respon"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Controller struct {
	service socialmedia.ISocialMediaService
}

func NewSocialMediaController(service socialmedia.ISocialMediaService) *Controller {
	return &Controller{
		service,
	}
}

// Get all social media by user id
// User can access to show all social media by user id
// Get all social media godoc
// @Summary Get all social media user
// @Description Get all social media user
// @Tags Social Media
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Success 200 {array} models.SocialMedia
// @Failure 500 {object} errors.RestErr
// @Router /socialmedia [get]
func (handler *Controller) GetAllSosmed(c *gin.Context) {
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

// Get social media by user id
// User can access to show social media by user id
// Get social media godoc
// @Summary Get social media user
// @Description Get social media user
// @Tags Social Media
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param socialMediaId path int true "Social Media ID"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} errors.RestErr
// @Failure 404 {object} errors.RestErr
// @Router /socialmedia/{socialMediaId} [get]
func (handler *Controller) GetSosmed(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))

	// Get param socialMediaId
	socialmediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil || uint(socialmediaId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewBadRequestError("invalid socialmedia ID"))

		return
	}

	res, err := handler.service.GetByUserId(userID, uint(socialmediaId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.NewNotFoundError(fmt.Sprintf("Photo Data with id: %d not found\n", socialmediaId)))
		return
	}
	c.JSON(http.StatusOK, res)
}

// Create social media by user id
// User can create social media by user id
// Create social media godoc
// @Summary Create social media user
// @Description Create social media user
// @Tags Social Media
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param requestCreate body models.RequestSocialMedia true "Create social media user"
// @Success 201 {object} models.RequestSocialMedia
// @Failure 400 {object} errors.RestErr
// @Router /socialmedia [post]
func (handler *Controller) CreateSosmed(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))
	Photo := models.RequestSocialMedia{}

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

// Update social media by user id
// User can update social media by user id
// Update social media godoc
// @Summary Update social media user
// @Description Update social media user
// @Tags Social Media
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param socialMediaId path int true "Social Media ID"
// @Param requestUpdate body models.RequestSocialMedia true "Update Social Media user"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} errors.RestErr
// @Router /socialmedia/{socialMediaId} [put]
func (handler *Controller) UpdateSosmed(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))
	sosmed := models.RequestSocialMedia{}

	// Get param photo_id
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil || uint(socialMediaId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewBadRequestError("invalid photo ID"))

		return
	}
	if err := c.ShouldBindJSON(&sosmed); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	}

	if res, err := handler.service.Update(sosmed, uint(socialMediaId), userID); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	} else {
		c.JSON(http.StatusCreated, res)
		return
	}
}

// Delete social media by user id
// User can delete social media by user id
// Delete social media godoc
// @Summary Delete social media user
// @Description Delete social media user
// @Tags Social Media
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param socialMediaId path int true "Social Media ID"
// @Router /socialmedia/{socialMediaId} [delete]
func (handler *Controller) DeleteSosmed(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))

	// Get Param socialMediaId
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil || uint(socialMediaId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewBadRequestError("invalid photo ID"))
		return
	}

	if _, err := handler.service.Delete(uint(socialMediaId), userID); err != nil {
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerError("Failed to delete photo"))
		return
	} else {
		c.JSON(http.StatusOK, respon.ResponseDeleted{
			Message: "Data Photo deleted successfully",
		})
	}
}
