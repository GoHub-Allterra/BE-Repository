package services

import (
	"errors"
	"gohub/features/post/domain"
)

type postUsecase struct {
	postData domain.PostData
}

func New(pd domain.PostData) domain.PostUsecase {
	return &postUsecase{
		postData: pd,
	}
}

func (ps *postUsecase) AddPost(data domain.Post, token int) (int, error) {
	if data.Caption != "" {

		add, err := ps.postData.Insert(data, token)
		if err != nil || add == 0 {
			return -1, err
		} else {
			return 1, nil
		}
	} else {
		return -1, errors.New("all input data must be filled")
	}
}

func (ps *postUsecase) GetAllPosts() ([]domain.Post, error) {
	dataAll, err := ps.postData.GetAll()
	if err != nil {
		return nil, errors.New("failed get all data")
	} else if len(dataAll) == 0 {
		return nil, errors.New("data is still empty")
	} else {
		return dataAll, nil
	}
}