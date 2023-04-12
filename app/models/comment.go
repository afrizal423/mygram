package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message string `json:"message" form:"message" valid:"required~Message is required"`
	PhotoID int    `json:"photo_id" form:"photo_id"`
	Photo   *Photo
	UserID  uint
	User    *User
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if p.PhotoID == 0 {
		return errors.New("Photo ID is required")
	}

	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	p.PhotoID = 0
	
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
