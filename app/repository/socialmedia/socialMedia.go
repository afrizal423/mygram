package socialmedia

import (
	"log"

	"github.com/afrizal423/mygram/app/models"
	"gorm.io/gorm"
)

type SocialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *SocialMediaRepository {
	return &SocialMediaRepository{
		db,
	}
}

func (sm *SocialMediaRepository) Create(req models.RequestSocialMedia, userID uint) (models.SocialMedia, error) {
	socialMedia := models.SocialMedia{}
	socialMedia.Name = req.Name
	socialMedia.SocialMediaUrl = req.SocialMediaUrl
	socialMedia.UserID = userID
	if err := sm.db.Create(&socialMedia).Error; err != nil {
		log.Println("Error creating social media:", err)
		return socialMedia, err
	}
	return socialMedia, nil
}

func (sm *SocialMediaRepository) GetAllByUserId(userID uint) ([]models.SocialMedia, error) {
	socialMedias := []models.SocialMedia{}
	if err := sm.db.Where("user_id = ?", userID).Order("id DESC").Find(&socialMedias).Error; err != nil {
		log.Println("Error get all social media:", err)
		return socialMedias, err
	}
	return socialMedias, nil
}

func (sm *SocialMediaRepository) GetByUserId(userID uint, id uint) (models.SocialMedia, error) {
	socialMedias := models.SocialMedia{}
	if err := sm.db.Where("user_id = ? AND id = ?", userID, id).Take(&socialMedias).Error; err != nil {
		log.Println("Error get social media:", err)
		return socialMedias, err
	}
	return socialMedias, nil
}

func (sm *SocialMediaRepository) Update(socialMedia models.SocialMedia, id, userID uint) (models.SocialMedia, error) {
	if err := sm.db.Model(&socialMedia).Where("id = ? AND user_id = ?", id, userID).Updates(models.SocialMedia{Name: socialMedia.Name, SocialMediaUrl: socialMedia.SocialMediaUrl}).Error; err != nil {
		log.Printf("Error updating social media: %+v data doesn't match\n", err)
		return socialMedia, err
	}
	// ambil data
	sm.db.Where("id = ? AND user_id = ?", id, userID).Take(&socialMedia)
	return socialMedia, nil
}

func (sm *SocialMediaRepository) Delete(id, userID uint) (models.SocialMedia, error) {
	socialMedia := models.SocialMedia{}
	if err := sm.db.Where("id = ? AND user_id = ?", id, userID).Delete(&socialMedia).Error; err != nil {
		log.Printf("Error deleting social media: %+v data doesn't match\n", err)
		return socialMedia, err
	}
	return socialMedia, nil
}
