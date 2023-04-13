package socialmedia

import "github.com/afrizal423/mygram/app/models"

type ISocialMediaService interface {
	// create data sosmed
	Create(req models.RequestSocialMedia, userID uint) (models.SocialMedia, error)
	// get all data by user id
	GetAllByUserId(userID uint) ([]models.SocialMedia, error)
	// get detail data by user id
	GetByUserId(userID uint, id uint) (models.SocialMedia, error)
	// update data sosmed
	Update(req models.RequestSocialMedia, id, userID uint) (models.SocialMedia, error)
	// hapus data sosmed
	Delete(id, userID uint) (models.SocialMedia, error)
}

type ISocialMediaRepository interface {
	// create data sosmed
	Create(req models.RequestSocialMedia, userID uint) (models.SocialMedia, error)
	// get all data by user id
	GetAllByUserId(userID uint) ([]models.SocialMedia, error)
	// get detail data by user id
	GetByUserId(userID uint, id uint) (models.SocialMedia, error)
	// update data sosmed
	Update(socialMedia models.SocialMedia, id, userID uint) (models.SocialMedia, error)
	// hapus data sosmed
	Delete(id, userID uint) (models.SocialMedia, error)
}
