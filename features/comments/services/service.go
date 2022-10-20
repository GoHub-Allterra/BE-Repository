package services

import (
	"gohub/features/comments/domain"
)

type CommentUsecase struct {
	commentData domain.DataInterface
}

func New(cd domain.DataInterface) domain.ServiceInterface {
	return &CommentUsecase{
		commentData: cd,
	}
}

func (cu *CommentUsecase) Insert(data domain.Comments, param int) (domain.Comments, error) {
	data, err := cu.commentData.AddComment(data, param)
	return data, err
}

func (cu *CommentUsecase) DeleteId(param, token int) (int, error) {

	_, err := cu.commentData.DeleteComent(param, token)
	if err != nil {
		return -1, err
	}

	return 1, nil
}
