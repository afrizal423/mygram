package comment

import (
	"log"

	"github.com/afrizal423/mygram/app/models"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db,
	}
}

func (cm *CommentRepository) GetAllByUserId(userID uint) ([]models.Comment, error) {
	komen := []models.Comment{}
	if err := cm.db.Where("user_id = ?", userID).Order("id DESC").Find(&komen).Error; err != nil {
		log.Println("Error get all photo:", err)
		return komen, err
	}
	return komen, nil
}

func (cm *CommentRepository) GetByUserId(userID uint, id uint) (models.Comment, error) {
	komen := models.Comment{}
	if err := cm.db.Where("user_id = ? AND id = ?", userID, id).Take(&komen).Error; err != nil {
		log.Println("Error get detail photo:", err)
		return komen, err
	}
	return komen, nil
}

func (cm *CommentRepository) Create(req models.Comment, userID uint) (models.Comment, error) {
	komen := models.Comment{}
	komen.Message = req.Message
	komen.PhotoID = req.PhotoID
	komen.UserID = userID
	if err := cm.db.Create(&komen).Error; err != nil {
		log.Println("Error creating komen:", err)
		return komen, err
	}
	return komen, nil
}

func (cm *CommentRepository) Update(komen models.Comment, id, userID uint) (models.Comment, error) {
	if err := cm.db.Model(&komen).Where("id = ? AND user_id = ?", id, userID).Updates(models.Comment{Message: komen.Message}).Error; err != nil {
		log.Printf("Error updating photo: %+v \n", err)
		return komen, err
	}
	// ambil data setelah update
	cm.db.Where("id = ? AND user_id = ?", id, userID).Take(&komen)
	return komen, nil
}

func (cm *CommentRepository) Delete(id, userID uint) (models.Comment, error) {
	komen := models.Comment{}
	if err := cm.db.Where("id = ? AND user_id = ? ", id, userID).Delete(&komen).Error; err != nil {
		log.Printf("Error deleting komen: %+v \n", err)
		return komen, err
	}
	return komen, nil
}

func (bk *CommentRepository) HitungFoto(id int) int64 {
	var count int64
	if err := bk.db.Model(&models.Photo{}).Where("id = ?", id).Count(&count).Error; err != nil {
		log.Println(err)
	}
	return count
}
