package services

import (
	"errors"
	"gohub/features/post/domain"
	mocks "gohub/mocks/features/post/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func TetstAddPost(t *testing.T) {
// 	repo := mocks.NewRepository(t)
// 	t.Run("sukses posting", func(t *testing.T) {
// 		repo.On("")
// 	})
// }

func TestPosts(t *testing.T) {
	postMock := new(mocks.PostData)
	input := domain.Post{ID: 1, User_ID: 15, Caption: "kaga rapih cees", Images: "images-img.jpg"}
	token := 1

	t.Run("create success", func(t *testing.T) {

		postMock.On("Insert", input, token).Return(1, nil).Once()

		useCase := New(postMock)
		res, _ := useCase.AddPost(input, token)
		assert.Equal(t, 1, res)
		postMock.AssertExpectations(t)
	})

	t.Run("create failed", func(t *testing.T) {

		postMock.On("Insert", mock.Anything, mock.Anything).Return(-1, errors.New("error")).Once()

		useCase := New(postMock)
		res, err := useCase.AddPost(input, token)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		postMock.AssertExpectations(t)

	})

}

func TestGetAll(t *testing.T) {
	getPosts := new(mocks.PostData)
	returnData := []domain.Post{{User_ID: 1, Caption: "Haloo", Images: "selfie.jpg"}}
	t.Run("Get All Success", func(t *testing.T) {

		getPosts.On("GetAll").Return(returnData, nil).Once()

		useCase := New(getPosts)
		res, err := useCase.GetAllPosts()
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		getPosts.AssertExpectations(t)

	})

	t.Run("Get All Failed", func(t *testing.T) {
		getPosts.On("GetAll").Return(nil, nil).Once()

		useCase := New(getPosts)
		res, _ := useCase.GetAllPosts()
		assert.Equal(t, 0, len(res))
		getPosts.AssertExpectations(t)

	})

	t.Run("Get All Failed", func(t *testing.T) {
		getPosts.On("GetAll").Return(nil, errors.New("error")).Once()

		useCase := New(getPosts)
		_, err := useCase.GetAllPosts()
		assert.Error(t, err)
		getPosts.AssertExpectations(t)

	})
}
func TestGetById(t *testing.T) {

	postMock := new(mocks.PostData)
	returnData := domain.Post{ID: 1, User_ID: 1, Caption: "Wakwaaww", Images: "sony-wakwaw.jpg"}
	param := 1

	t.Run("Get by id success", func(t *testing.T) {
		postMock.On("GetById", param).Return(returnData, nil).Once()

		useCase := New(postMock)
		res, _ := useCase.SelectById(param)
		assert.Equal(t, param, int(res.ID))
		postMock.AssertExpectations(t)

	})

	t.Run("Get by id failed", func(t *testing.T) {

		postMock.On("GetById", param).Return(domain.Post{}, errors.New("error")).Once()

		useCase := New(postMock)
		param := 1
		res, err := useCase.SelectById(param)
		assert.Error(t, err)
		assert.NotEqual(t, param, int(res.ID))
		postMock.AssertExpectations(t)

	})

}
func TestPut(t *testing.T) {

	postMock := new(mocks.PostData)
	input := domain.Post{User_ID: 1, Caption: "Halaaaawwww", Images: "say-hello.PNG"}
	param := 1
	token := 1

	t.Run("update succes", func(t *testing.T) {

		postMock.On("PutPost", param, token, input).Return(1, nil).Once()

		useCase := New(postMock)
		res, _ := useCase.UpdatePost(param, token, input)
		assert.Equal(t, 1, res)
		postMock.AssertExpectations(t)

	})

	t.Run("update failed", func(t *testing.T) {

		postMock.On("PutPost", param, token, input).Return(-1, errors.New("error")).Once()

		useCase := New(postMock)
		res, err := useCase.UpdatePost(param, token, input)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		postMock.AssertExpectations(t)

	})

}

func TestDelete(t *testing.T) {

	postMock := new(mocks.PostData)
	token := 1
	param := 3

	t.Run("delete succes", func(t *testing.T) {

		postMock.On("DeletedId", param, token).Return(1, nil).Once()

		useCase := New(postMock)
		res, _ := useCase.DeletedPost(param, token)
		assert.Equal(t, 1, res)
		postMock.AssertExpectations(t)

	})

	t.Run("delete failed", func(t *testing.T) {

		postMock.On("DeletedId", param, token).Return(-1, errors.New("error")).Once()

		useCase := New(postMock)
		res, err := useCase.DeletedPost(param, token)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		postMock.AssertExpectations(t)

	})

}

func TestGetMyPosts(t *testing.T) {
	postMock := new(mocks.PostData)
	returnData := ([]domain.Post{{ID: 1, User_ID: 1, Caption: "Wakwaaww", Images: "sony-wakwaw.jpg", Created_At: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), Updated_At: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), Comment: []domain.CommentsCore(nil)}})
	var param uint
	param = 3
	t.Run("success get my posts", func(t *testing.T) {
		postMock.On("GetAllPostsByID", param).Return(returnData, nil).Once()

		useCase := New(postMock)
		res, _ := useCase.GetMyPosts(param)
		assert.Equal(t, res, res)
		postMock.AssertExpectations(t)

	})
}
