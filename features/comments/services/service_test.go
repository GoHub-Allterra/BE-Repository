package services

import (
	"errors"
	"gohub/features/comments/domain"
	mocks "gohub/mocks/features/comments/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAdd(t *testing.T) {
	commentMock := new(mocks.DataInterface)
	input := domain.Comments{User_ID: 1, Comment: "up keras gan"}
	// param := 1
	// token := 1

	t.Run("create success", func(t *testing.T) {

		commentMock.On("AddComment", mock.Anything).Return(input, nil).Once()

		useCase := New(commentMock)
		res, err := useCase.Insert(input)
		assert.NotEqual(t, res, err)
		commentMock.AssertExpectations(t)
	})

	t.Run("create failed", func(t *testing.T) {

		commentMock.On("AddComment", mock.Anything).Return(domain.Comments{}, errors.New("error")).Once()

		useCase := New(commentMock)
		_, err := useCase.Insert(input)
		assert.Equal(t, -1, -1)
		assert.Error(t, err)
		commentMock.AssertExpectations(t)

	})

}
func TestDelete(t *testing.T) {

	commentMock := new(mocks.DataInterface)
	token := 1
	param := 3

	t.Run("delete succes", func(t *testing.T) {

		commentMock.On("DeleteComent", param, token).Return(1, nil).Once()

		useCase := New(commentMock)
		res, _ := useCase.DeleteId(param, token)
		assert.Equal(t, 1, res)
		commentMock.AssertExpectations(t)

	})

	t.Run("delete failed", func(t *testing.T) {

		commentMock.On("DeleteComent", param, token).Return(-1, errors.New("error")).Once()

		useCase := New(commentMock)
		res, err := useCase.DeleteId(param, token)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		commentMock.AssertExpectations(t)

	})

}
