package photo

import (
	"log"

	"github.com/afrizal423/mygram/app/models"
	"gorm.io/gorm"
)

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{
		db,
	}
}

func (p *PhotoRepository) GetAllByUserId(userID uint) ([]models.Photo, error) {
	photos := []models.Photo{}
	if err := p.db.Where("user_id = ?", userID).Order("id DESC").Find(&photos).Error; err != nil {
		log.Println("Error get all photo:", err)
		return photos, err
	}
	return photos, nil
}

func (p *PhotoRepository) GetByUserId(userID uint, id uint) (models.Photo, error) {
	photos := models.Photo{}
	if err := p.db.Where("user_id = ? AND id = ?", userID, id).Take(&photos).Error; err != nil {
		log.Println("Error get detail photo:", err)
		return photos, err
	}
	return photos, nil
}

func (p *PhotoRepository) Create(req models.Photo, userID uint) (models.Photo, error) {
	photo := models.Photo{}
	photo.Title = req.Title
	photo.PhotoUrl = req.PhotoUrl
	photo.Caption = req.Caption
	photo.UserID = userID
	if err := p.db.Create(&photo).Error; err != nil {
		log.Println("Error creating photo:", err)
		return photo, err
	}
	return photo, nil
}

func (p *PhotoRepository) Update(photo models.Photo, id, userID uint) (models.Photo, error) {
	if err := p.db.Model(&photo).Where("id = ? AND user_id = ?", id, userID).Updates(models.Photo{Title: photo.Title, Caption: photo.Caption, PhotoUrl: photo.PhotoUrl}).Error; err != nil {
		log.Printf("Error updating photo: %+v \n", err)
		return photo, err
	}
	// ambil data setelah update
	p.db.Where("id = ? AND user_id = ?", id, userID).Take(&photo)
	return photo, nil
}

func (p *PhotoRepository) Delete(id, userID uint) (models.Photo, error) {
	photo := models.Photo{}
	if err := p.db.Where("id = ? AND user_id = ?", id, userID).Delete(&photo).Error; err != nil {
		log.Printf("Error deleting photo: %+v \n", err)
		return photo, err
	}
	return photo, nil
}
