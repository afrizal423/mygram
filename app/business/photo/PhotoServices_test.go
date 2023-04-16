package photo_test

import (
	"testing"

	"github.com/afrizal423/mygram/app/business/photo"
	"github.com/afrizal423/mygram/app/business/photo/mocks"
	"github.com/afrizal423/mygram/app/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var photoRepoInterfaceMock = &mocks.IPhotoRepository{Mock: mock.Mock{}}
var photoUseCaseInterface = photo.PhotoService{Repository: photoRepoInterfaceMock}

func TestInsertPhoto(t *testing.T) {
	photoDummy := models.Photo{
		Title:    "ini judul",
		Caption:  "ini caption",
		PhotoUrl: "google.com/logo.ico",
		UserID:   1,
	}
	photoRepoInterfaceMock.On("Create", mock.AnythingOfType("models.Photo"), uint(1)).Return(photoDummy, nil).Once()
	photoInsert := models.Photo{
		Title:    "ini judul",
		Caption:  "ini caption",
		PhotoUrl: "google.com/logo.ico",
		UserID:   1,
	}
	data, err := photoUseCaseInterface.Create(photoInsert, 1)
	assert.Equal(t, nil, err)
	assert.Equal(t, photoDummy, data)
}

func TestGetDataPhoto(t *testing.T) {

	t.Run("Get All data", func(t *testing.T) {
		photoDummy := []models.Photo{
			{
				Title:    "ini judul",
				Caption:  "ini caption",
				PhotoUrl: "google.com/logo.ico",
				UserID:   1,
			},
		}
		photoRepoInterfaceMock.On("GetAllByUserId", uint(1)).Return(photoDummy, nil).Once()

		data, err := photoUseCaseInterface.GetAllByUserId(1)
		assert.Equal(t, nil, err)
		assert.Equal(t, photoDummy, data)

	})

	t.Run("Get Single data", func(t *testing.T) {
		photoDummy := models.Photo{
			Title:    "ini judul",
			Caption:  "ini caption",
			PhotoUrl: "google.com/logo.ico",
			UserID:   1,
		}
		photoRepoInterfaceMock.On("GetByUserId", uint(1), uint(1)).Return(photoDummy, nil).Once()

		data, err := photoUseCaseInterface.GetByUserId(uint(1), uint(1))
		assert.Equal(t, nil, err)
		assert.Equal(t, photoDummy, data)
	})
}

func TestUpdateDataPhoto(t *testing.T) {
	var DummyEdit models.Photo
	photoRepoInterfaceMock.On("Update", mock.AnythingOfType("models.Photo"), mock.Anything, mock.Anything).Return(models.Photo{}, nil).Once()
	dataSudahUpdate := models.Photo{
		Title:    "ini judul",
		Caption:  "ini caption",
		PhotoUrl: "google.com/logo.ico",
		UserID:   1,
	}

	data, err := photoUseCaseInterface.Update(dataSudahUpdate, uint(1), uint(1))
	// fmt.Println(data)
	assert.Equal(t, nil, err)
	assert.Equal(t, DummyEdit, data)
}

func TestDeleteDataPhoto(t *testing.T) {
	photoDummy := models.Photo{
		GormModel: models.GormModel{
			ID: 1,
		},
		Title:    "ini judul",
		Caption:  "ini caption",
		PhotoUrl: "google.com/logo.ico",
		UserID:   1,
	}
	photoRepoInterfaceMock.On("Delete", mock.Anything, mock.Anything).Return(photoDummy, nil).Once()
	hps, err := photoUseCaseInterface.Delete(uint(1), uint(1))
	assert.Equal(t, nil, err)
	assert.Equal(t, photoDummy, hps)
}
