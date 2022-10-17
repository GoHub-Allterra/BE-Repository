package services

import (
	"errors"
	"gohub/features/user/domain"
	"gohub/mocks/features/user/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Add User", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{ID: uint(1), Name: "Fatur", HP: "08123", Password: "fatur123"}, nil).Once()
		srv := New(repo)
		input := domain.Core{ID: uint(1), Name: "Fatur", HP: "08123", Password: "fatur123"}
		res, err := srv.AddUser(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Add User", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New("error add user")).Once()
		srv := New(repo)
		res, err := srv.AddUser(domain.Core{})
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Database error", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New("error add user")).Once()
		srv := New(repo)
		res, err := srv.AddUser(domain.Core{ID: 5, Name: "ian", HP: "08213"})
		assert.ErrorContains(t, err, "database")
		assert.Empty(t, res)
	})
}

func TestDeleteUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Delete User", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{ID: uint(1), Name: "Fatur", HP: "08123", Password: "fatur123"}, nil).Once()
		srv := New(repo)
		res, err := srv.DeleteUser(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Delete User", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{}, errors.New("error")).Once()
		srv := New(repo)
		res, err := srv.DeleteUser(1)
		assert.NotEmpty(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Get User", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(domain.Core{ID: uint(1), Name: "Fatur", HP: "08123", Password: "fatur123"}, nil).Once()
		srv := New(repo)
		res, err := srv.Get(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Get User", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(domain.Core{}, errors.New("error get id")).Once()
		srv := New(repo)
		res, err := srv.Get(1)
		assert.Empty(t, res, "seharusnya res ada isinya")
		assert.Nil(t, err, "seharusnya err itu nil")
		repo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Update User", func(t *testing.T) {
		repo.On("Edit", mock.Anything, mock.Anything).Return(domain.Core{ID: 1, Name: "fatur", HP: "08123", Password: "fatur123"}, nil).Once()
		srv := New(repo)
		input := domain.Core{ID: 1, Name: "fatur", HP: "08123", Password: "fatur123"}
		res, err := srv.UpdateUser(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Update User", func(t *testing.T) {
		repo.On("Edit", mock.Anything, mock.Anything).Return(domain.Core{}, errors.New("error update data")).Once()
		srv := New(repo)
		input := domain.Core{}
		res, err := srv.UpdateUser(input)
		assert.NotEmpty(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}