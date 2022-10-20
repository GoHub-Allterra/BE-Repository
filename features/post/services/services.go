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

func (ps *postUsecase) UpdatePost(param, token int, data domain.Post) (int, error) {
	row, err := ps.postData.PutPost(param, token, data)
	if err != nil || row == 0 {
		return -1, err
	}

	return 1, nil
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

func (ps *postUsecase) GetMyPosts(id uint) ([]domain.Post, error) {
	posts, err := ps.postData.GetAllPostsByID(id)
	return posts, err
}

func (ps *postUsecase) SelectById(param int) (domain.Post, error) {

	dataId, err := ps.postData.GetById(param)
	if err != nil {
		return domain.Post{}, err
	}

	return dataId, nil

}

func (ps *postUsecase) DeletedPost(param, token int) (int, error) {
	row, err := ps.postData.DeletedId(param, token)
	if err != nil || row == 0 {
		return -1, err
	}

	return int(row), nil
}
