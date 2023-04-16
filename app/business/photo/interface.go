package photo

import "github.com/afrizal423/mygram/app/models"

type IPhotoService interface {
	// get data foto by id
	GetAllByUserId(userID uint) ([]models.Photo, error)
	// get single data by user id
	GetByUserId(userID uint, id uint) (models.Photo, error)
	// create data foto
	Create(req models.Photo, userID uint) (models.Photo, error)
	// update data foto
	Update(req models.Photo, id, userID uint) (models.Photo, error)
	// delete data foto
	Delete(id, userID uint) (models.Photo, error)
}

//go:generate mockery --name IPhotoRepository
type IPhotoRepository interface {
	// get data foto by id
	GetAllByUserId(userID uint) ([]models.Photo, error)
	// get single data by user id
	GetByUserId(userID uint, id uint) (models.Photo, error)
	// create data foto
	Create(req models.Photo, userID uint) (models.Photo, error)
	// update data foto
	Update(req models.Photo, id, userID uint) (models.Photo, error)
	// delete data foto
	Delete(id, userID uint) (models.Photo, error)
}
