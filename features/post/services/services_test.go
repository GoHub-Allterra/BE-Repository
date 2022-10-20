package services

import (
	mocks "gohub/mocks/features/user/domain"
	"testing"
)

func TetstAddPost(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("sukses posting", func(t *testing.T) {
		repo.On("")
	})
}

func TestGetCommentByIdPosts(t *testing.T) {

}

func TestGetAll(t *testing.T) {

}
