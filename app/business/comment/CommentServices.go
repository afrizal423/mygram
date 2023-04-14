package comment

import (
	"errors"
	"log"

	"github.com/afrizal423/mygram/app/models"
)

type CommentService struct {
	repository ICommentRepository
}

func NewCommentService(repository ICommentRepository) *CommentService {
	return &CommentService{
		repository,
	}
}

func (s *CommentService) GetAllByUserId(userID uint) ([]models.Comment, error) {

	if photos, err := s.repository.GetAllByUserId(userID); err != nil {
		log.Println("Data not found")
		return photos, err
	} else {
		return photos, nil
	}
}

func (s *CommentService) GetByUserId(userID uint, id uint) (models.Comment, error) {

	if photo, err := s.repository.GetByUserId(userID, id); err != nil {
		log.Println("Data not found")
		return photo, err
	} else {
		return photo, nil
	}
}

func (s *CommentService) Create(req models.RequestComments, userID uint) (models.Comment, error) {
	komen := models.Comment{}
	komen.Message = req.Message
	komen.PhotoID = req.PhotoID
	komen.UserID = userID
	if s.repository.HitungFoto(req.PhotoID) == 0 {
		return komen, errors.New("data foto kosong")
	}
	if photo, err := s.repository.Create(komen, userID); err != nil {
		log.Println("Failed to create photo")
		return photo, err
	} else {
		return photo, nil
	}
}

func (s *CommentService) Update(req models.RequestComments, id, userID, photoID uint) (models.Comment, error) {
	komen := models.Comment{}
	komen.Message = req.Message
	komen.PhotoID = req.PhotoID
	komen.UserID = userID
	if photo, err := s.repository.Update(komen, id, userID); err != nil {
		log.Println("Failed to update photo")
		return photo, err
	} else {
		return photo, nil
	}
}

func (s *CommentService) Delete(id, userID uint) (models.Comment, error) {

	if photo, err := s.repository.Delete(id, userID); err != nil {
		log.Println("Failed to delete photo")
		return photo, err
	} else {
		return photo, nil
	}
}
