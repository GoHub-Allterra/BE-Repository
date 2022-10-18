package repository

import (
	"errors"
	"gohub/features/post/domain"

	"gorm.io/gorm"
)

type postData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.PostData {
	return &postData{
		db: DB,
	}
}

func (pd *postData) Insert(data domain.Post, token int) (int, error) {
	var datacheck domain.User
	txcheck := pd.db.Where("ID=?", token).First(&datacheck)
	if txcheck.Error != nil {
		return -1, errors.New("error tx")
	}

	if int(datacheck.ID) != token {
		return -1, errors.New("not have access")
	}

	dataModel := ToEntity(data)
	dataModel.User_ID = uint(token)
	tx := pd.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil

}

func (pd *postData) GetAll() ([]domain.Post, error) {
	var dataPost []Post
	tx := pd.db.Find(&dataPost)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataPostUser := toPostList(dataPost)

	return dataPostUser, nil
}

func (pd *postData) GetById(param int) (domain.Post, error) {
	var dataId Post
	tx := pd.db.First(&dataId, param)
	if tx.Error != nil {
		return domain.Post{}, tx.Error
	}

	postId := dataId.toPostUser()
	return postId, nil
}
