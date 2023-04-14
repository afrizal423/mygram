package comment

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/afrizal423/mygram/app/business/comment"
	"github.com/afrizal423/mygram/app/models"
	"github.com/afrizal423/mygram/pkg/utils/errors"
	"github.com/afrizal423/mygram/pkg/utils/respon"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Controller struct {
	service comment.ICommentService
}

func NewCommentController(service comment.ICommentService) *Controller {
	return &Controller{
		service,
	}
}

// Get all Comment by user id
// User can access to show all Comment by user id
// Get all Comment godoc
// @Summary Get all Comment user
// @Description Get all Comment user
// @Tags Comment
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Success 200 {array} models.Comment
// @Failure 500 {object} errors.RestErr
// @Router /comment [get]
func (handler *Controller) GetAllComment(c *gin.Context) {
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

// Get Comment by user id
// User can access to show Comment by user id
// Get Comment godoc
// @Summary Get Comment user
// @Description Get Comment user
// @Tags Comment
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param commentId path int true "Comment ID"
// @Success 200 {object} models.Comment
// @Failure 400 {object} errors.RestErr
// @Failure 404 {object} errors.RestErr
// @Router /comment/{commentId} [get]
func (handler *Controller) GetComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))

	// Get param commentId
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil || uint(commentId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewBadRequestError("invalid socialmedia ID"))

		return
	}

	res, err := handler.service.GetByUserId(userID, uint(commentId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.NewNotFoundError(fmt.Sprintf("Photo Data with id: %d not found\n", commentId)))
		return
	}
	c.JSON(http.StatusOK, res)
}

// Create Comment by user id
// User can create Comment by user id
// Create Comment godoc
// @Summary Create Comment user
// @Description Create Comment user
// @Tags Comment
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param requestCreate body models.RequestComments true "Create Comment user"
// @Success 201 {object} models.RequestComments
// @Failure 400 {object} errors.RestErr
// @Router /comment [post]
func (handler *Controller) CreateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))
	Photo := models.RequestComments{}

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

// Update Comment by user id
// User can update Comment by user id
// Update Comment godoc
// @Summary Update Comment user
// @Description Update Comment user
// @Tags Comment
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param commentId path int true "Comment ID"
// @Param requestUpdate body models.RequestComments true "Update Comment user"
// @Success 200 {object} models.Comment
// @Failure 400 {object} errors.RestErr
// @Router /comment/{commentId} [put]
func (handler *Controller) UpdateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))
	sosmed := models.RequestComments{}

	// Get param photo_id
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil || uint(commentId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewBadRequestError("invalid photo ID"))

		return
	}
	if err := c.ShouldBindJSON(&sosmed); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	}

	if res, err := handler.service.Update(sosmed, uint(commentId), userID, uint(sosmed.PhotoID)); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	} else {
		c.JSON(http.StatusCreated, res)
		return
	}
}

// Delete Comment by user id
// User can delete Comment by user id
// Delete Comment godoc
// @Summary Delete Comment user
// @Description Delete Comment user
// @Tags Comment
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param commentId path int true "Comment ID"
// @Router /comment/{commentId} [delete]
func (handler *Controller) DeleteComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))

	// Get Param commentId
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil || uint(commentId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewBadRequestError("invalid photo ID"))
		return
	}

	if _, err := handler.service.Delete(uint(commentId), userID); err != nil {
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerError("Failed to delete photo"))
		return
	} else {
		c.JSON(http.StatusOK, respon.ResponseDeleted{
			Message: "Data Photo deleted successfully",
		})
	}
}
