package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `json:"title" form:"title" valid:"required~Title of your photo is required"`
	Caption  string `json:"caption" form:"caption" valid:"required~Caption of your photo is required"`
	PhotoUrl string `json:"photo_url" form:"photo_url" valid:"required~Photo URL of your photo is required"`
	UserID   uint
	User     *User
}

// create, update
type RequestPhoto struct {
	Title    string `gorm:"-:all" json:"title" form:"title" binding:"required"`
	Caption  string `gorm:"-:all" json:"caption,omitempty" form:"caption,omitempty"`
	PhotoURL string `gorm:"-:all" json:"photo_url" form:"photo_url" binding:"required,url"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
