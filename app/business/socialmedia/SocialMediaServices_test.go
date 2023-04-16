package socialmedia_test

import (
	"testing"

	"github.com/afrizal423/mygram/app/business/socialmedia"
	"github.com/afrizal423/mygram/app/business/socialmedia/mocks"
	"github.com/afrizal423/mygram/app/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var sosmedRepoInterfaceMock = &mocks.ISocialMediaRepository{Mock: mock.Mock{}}
var sosmedUseCaseInterface = socialmedia.SocialMediaService{Repository: sosmedRepoInterfaceMock}

func TestInsertSosmed(t *testing.T) {
	sosmdedDummy := models.SocialMedia{
		Name:           "sosmed",
		SocialMediaUrl: "google.com",
		UserID:         1,
	}
	sosmedRepoInterfaceMock.On("Create", mock.AnythingOfType("models.RequestSocialMedia"), uint(1)).Return(sosmdedDummy, nil).Once()
	sosmdedInsert := models.RequestSocialMedia{
		Name:           "sosmed",
		SocialMediaUrl: "google.com",
	}
	data, err := sosmedUseCaseInterface.Create(sosmdedInsert, 1)
	assert.Equal(t, nil, err)
	assert.Equal(t, sosmdedDummy, data)
}

func TestGetDataSosmed(t *testing.T) {

	t.Run("Get All data", func(t *testing.T) {
		sosmdedDummy := []models.SocialMedia{
			{
				Name:           "sosmed",
				SocialMediaUrl: "google.com",
				UserID:         1,
			},
		}
		sosmedRepoInterfaceMock.On("GetAllByUserId", uint(1)).Return(sosmdedDummy, nil).Once()

		data, err := sosmedUseCaseInterface.GetAllByUserId(1)
		assert.Equal(t, nil, err)
		assert.Equal(t, sosmdedDummy, data)

	})

	t.Run("Get Single data", func(t *testing.T) {
		sosmdedDummy := models.SocialMedia{
			Name:           "sosmed",
			SocialMediaUrl: "google.com",
			UserID:         1,
		}
		sosmedRepoInterfaceMock.On("GetByUserId", uint(1), uint(1)).Return(sosmdedDummy, nil).Once()

		data, err := sosmedUseCaseInterface.GetByUserId(uint(1), uint(1))
		assert.Equal(t, nil, err)
		assert.Equal(t, sosmdedDummy, data)
	})
}

func TestUpdateDataSosmed(t *testing.T) {
	var DummyEdit models.SocialMedia
	sosmedRepoInterfaceMock.On("Update", mock.AnythingOfType("models.SocialMedia"), mock.Anything, mock.Anything).Return(models.SocialMedia{}, nil).Once()
	dataSudahUpdate := models.RequestSocialMedia{
		Name:           "sosmed",
		SocialMediaUrl: "google.com",
	}

	data, err := sosmedUseCaseInterface.Update(dataSudahUpdate, uint(1), uint(1))
	// fmt.Println(data)
	assert.Equal(t, nil, err)
	assert.Equal(t, DummyEdit, data)
}

func TestDeleteDataSosmed(t *testing.T) {
	sosmdedDummy := models.SocialMedia{
		GormModel: models.GormModel{
			ID: 1,
		},
		Name:           "sosmed",
		SocialMediaUrl: "google.com",
		UserID:         1,
	}
	sosmedRepoInterfaceMock.On("Delete", mock.Anything, mock.Anything).Return(sosmdedDummy, nil).Once()
	hps, err := sosmedUseCaseInterface.Delete(uint(1), uint(1))
	assert.Equal(t, nil, err)
	assert.Equal(t, sosmdedDummy, hps)
}
