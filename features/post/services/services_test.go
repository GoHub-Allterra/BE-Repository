package services

import (
	mocks "gohub/mocks/features/user/domain"
	"testing"
)


func AddPost(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("sukses posting", func(t *testing.T) {
		repo.On("")
	})
}