package photo

import (
	"log"

	"github.com/afrizal423/mygram/app/models"
)

type PhotoService struct {
	Repository IPhotoRepository
}

func NewPhotoService(repository IPhotoRepository) *PhotoService {
	return &PhotoService{
		repository,
	}
}

func (s *PhotoService) GetAllByUserId(userID uint) ([]models.Photo, error) {

	if photos, err := s.Repository.GetAllByUserId(userID); err != nil {
		log.Println("Data not found")
		return photos, err
	} else {
		return photos, nil
	}
}

func (s *PhotoService) GetByUserId(userID uint, id uint) (models.Photo, error) {

	if photo, err := s.Repository.GetByUserId(userID, id); err != nil {
		log.Println("Data not found")
		return photo, err
	} else {
		return photo, nil
	}
}

func (s *PhotoService) Create(req models.Photo, userID uint) (models.Photo, error) {
	if photo, err := s.Repository.Create(req, userID); err != nil {
		log.Println("Failed to create photo")
		return photo, err
	} else {
		return photo, nil
	}
}

func (s *PhotoService) Update(req models.Photo, id, userID uint) (models.Photo, error) {

	if photo, err := s.Repository.Update(req, id, userID); err != nil {
		log.Println("Failed to update photo")
		return photo, err
	} else {
		return photo, nil
	}
}

func (s *PhotoService) Delete(id, userID uint) (models.Photo, error) {

	if photo, err := s.Repository.Delete(id, userID); err != nil {
		log.Println("Failed to delete photo")
		return photo, err
	} else {
		return photo, nil
	}
}
