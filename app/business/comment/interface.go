package comment

import "github.com/afrizal423/mygram/app/models"

type ICommentService interface {
	// get data komentar by id
	GetAllByUserId(userID uint) ([]models.Comment, error)
	// get single data by user id
	GetByUserId(userID uint, id uint) (models.Comment, error)
	// create data komentar
	Create(req models.RequestComments, userID uint) (models.Comment, error)
	// update data komentar
	Update(req models.RequestComments, id, userID, photoID uint) (models.Comment, error)
	// delete data komentar
	Delete(id, userID uint) (models.Comment, error)
}

//go:generate mockery --name ICommentRepository
type ICommentRepository interface {
	// get data komentar by id
	GetAllByUserId(userID uint) ([]models.Comment, error)
	// get single data by user id
	GetByUserId(userID uint, id uint) (models.Comment, error)
	// create data komentar
	Create(req models.Comment, userID uint) (models.Comment, error)
	// update data komentar
	Update(req models.Comment, id, userID uint) (models.Comment, error)
	// delete data komentar
	Delete(id, userID uint) (models.Comment, error)

	HitungFoto(id int) int64
}
