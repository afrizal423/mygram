package database

import (
	"github.com/afrizal423/mygram/app/models"
	"gorm.io/gorm"
)

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
}
