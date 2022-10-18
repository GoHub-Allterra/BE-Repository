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

// func (ps *postUsecase) GetAllPosts() ([]domain.Post, []domain.User, [][]string, error) {
// 	data, userdata, post_images, err := ps.postData.GetAll()
// 	return data, userdata, post_images, err
// }

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

// func (ps *postUsecase) AddPostImages(images []string, postID uint) error {
// 	err := ps.postData.InsertPostImages(images, postID)
// 	return err
// }

// func (ps *postUsecase) GetMyPosts(id uint) ([]domain.Post, domain.User, [][]string, error) {
// 	posts, userdata, postimages, err := ps.postData.GetAllPostsByID(id)
// 	return posts, userdata, postimages, err
// }

// func (ps *postUsecase) GetSpecificPost(id uint) (domain.Post, domain.User, []string, []domain.Comment, []domain.User, error) {
// 	post, userdata, postimages, comments, commentUserData, err := ps.postData.GetPostByID(id)
// 	return post, userdata, postimages, comments, commentUserData, err
// }

// func (ps *postUsecase) UpdatePost(id uint, updateData domain.Post) (domain.Post, error) {
// 	data, err := ps.postData.Update(id, updateData)
// 	return data, err
// }

// func (ps *postUsecase) DeletePost(id uint, userID uint) error {
// 	err := ps.postData.Delete(id, userID)
// 	return err
// }
