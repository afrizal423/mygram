package socialmedia

import (
	"log"

	"github.com/afrizal423/mygram/app/models"
)

type SocialMediaService struct {
	Repository ISocialMediaRepository
}

func NewSocialMediaService(repository ISocialMediaRepository) *SocialMediaService {
	return &SocialMediaService{
		repository,
	}
}

func (s *SocialMediaService) Create(req models.RequestSocialMedia, userID uint) (models.SocialMedia, error) {
	if socialMedia, err := s.Repository.Create(req, userID); err != nil {
		log.Println("Failed to create socialMedia")
		return socialMedia, err
	} else {
		return socialMedia, nil
	}
}

func (s *SocialMediaService) GetAllByUserId(userID uint) ([]models.SocialMedia, error) {

	if socialMedias, err := s.Repository.GetAllByUserId(userID); err != nil {
		log.Println("Data not found")
		return socialMedias, err
	} else {
		return socialMedias, nil
	}
}

func (s *SocialMediaService) GetByUserId(userID uint, id uint) (models.SocialMedia, error) {
	if socialMedia, err := s.Repository.GetByUserId(userID, id); err != nil {
		log.Println("Data not found")
		return socialMedia, err
	} else {
		return socialMedia, nil
	}
}

func (s *SocialMediaService) Update(req models.RequestSocialMedia, id, userID uint) (models.SocialMedia, error) {
	socialMedia := models.SocialMedia{}
	socialMedia.Name = req.Name
	socialMedia.SocialMediaUrl = req.SocialMediaUrl
	socialMedia.UserID = userID
	if socialMedia, err := s.Repository.Update(socialMedia, id, userID); err != nil {
		log.Println("Failed to update socialMedia")
		return socialMedia, err
	} else {
		return socialMedia, nil
	}
}

func (s *SocialMediaService) Delete(id, userID uint) (models.SocialMedia, error) {
	if socialMedia, err := s.Repository.Delete(id, userID); err != nil {
		log.Println("Failed to delete socialMedia")
		return socialMedia, err
	} else {
		return socialMedia, nil
	}
}
