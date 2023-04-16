package comment_test

import (
	"testing"

	"github.com/afrizal423/mygram/app/business/comment"
	"github.com/afrizal423/mygram/app/business/comment/mocks"
	"github.com/afrizal423/mygram/app/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var CommentRepoInterfaceMock = &mocks.ICommentRepository{Mock: mock.Mock{}}
var CommentUseCaseInterface = comment.CommentService{Repository: CommentRepoInterfaceMock}

func TestInsertComment(t *testing.T) {
	commentDummy := models.Comment{
		Message: "ini komentar",
		PhotoID: 1,
		UserID:  1,
	}
	CommentRepoInterfaceMock.On("HitungFoto", commentDummy.PhotoID).Return(int64(1)).Once()
	CommentRepoInterfaceMock.On("Create", mock.AnythingOfType("models.Comment"), uint(1)).Return(commentDummy, nil).Once()
	commentINsert := models.RequestComments{
		Message: "ini komentar",
		PhotoID: 1,
	}
	data, err := CommentUseCaseInterface.Create(commentINsert, 1)
	assert.Equal(t, nil, err)
	assert.Equal(t, commentDummy, data)
}

func TestGetDataComment(t *testing.T) {

	t.Run("Get All data", func(t *testing.T) {
		commentDummy := []models.Comment{
			{
				Message: "ini komentar",
				PhotoID: 1,
				UserID:  1,
			},
		}
		CommentRepoInterfaceMock.On("GetAllByUserId", uint(1)).Return(commentDummy, nil).Once()

		data, err := CommentUseCaseInterface.GetAllByUserId(1)
		assert.Equal(t, nil, err)
		assert.Equal(t, commentDummy, data)

	})

	t.Run("Get Single data", func(t *testing.T) {
		commentDummy := models.Comment{
			Message: "ini komentar",
			PhotoID: 1,
			UserID:  1,
		}
		CommentRepoInterfaceMock.On("GetByUserId", uint(1), uint(1)).Return(commentDummy, nil).Once()

		data, err := CommentUseCaseInterface.GetByUserId(uint(1), uint(1))
		assert.Equal(t, nil, err)
		assert.Equal(t, commentDummy, data)
	})
}

func TestUpdateDataComment(t *testing.T) {
	var DummyEdit models.Comment
	CommentRepoInterfaceMock.On("Update", mock.AnythingOfType("models.Comment"), mock.Anything, mock.Anything, mock.Anything).Return(models.Comment{}, nil).Once()
	dataSudahUpdate := models.RequestComments{
		Message: "ini komentar",
		PhotoID: 1,
	}

	data, err := CommentUseCaseInterface.Update(dataSudahUpdate, uint(1), uint(1), uint(1))
	// fmt.Println(data)
	assert.Equal(t, nil, err)
	assert.Equal(t, DummyEdit, data)
}

func TestDeleteDataComment(t *testing.T) {
	commentDummy := models.Comment{
		GormModel: models.GormModel{
			ID: 1,
		},
		Message: "ini komentar",
		PhotoID: 1,
		UserID:  1,
	}
	CommentRepoInterfaceMock.On("Delete", mock.Anything, mock.Anything).Return(commentDummy, nil).Once()
	hps, err := CommentUseCaseInterface.Delete(uint(1), uint(1))
	assert.Equal(t, nil, err)
	assert.Equal(t, commentDummy, hps)
}
