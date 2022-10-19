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

func (pd *postData) DeletedId(param, token int) (int, error) {

	dataCart := Post{}
	idCheck := pd.db.First(&dataCart, param)
	if idCheck.Error != nil {
		return 0, idCheck.Error
	}
	if uint(token) != dataCart.User_ID {
		return -1, errors.New("you don't have access")
	}
	result := pd.db.Delete(&Post{}, param)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to delete data")
	}
	return int(result.RowsAffected), nil

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

func (pd *postData) GetAllPostsByID(id uint) ([]domain.Post, error) {
	var postData []domain.Post
	pd.db.Where("user_id", id).Find(&postData)
	if len(postData) < 1 {
		return []domain.Post{}, errors.New("no postData found")
	}

	var userData domain.User
	pd.db.Raw("SELECT name FROM users WHERE id = ?", id).Scan(&userData)

	return postData, nil
}

func (pd *postData) PutPost(param, token int, dataUpdate domain.Post) (int, error) {
	var dataCheck Post
	tx := pd.db.First(&dataCheck, param)
	if tx.Error != nil {
		return -1, tx.Error
	}
	postId := dataCheck.toPostUser()

	if postId.User_ID == uint(token) {
		var data Post
		data.Caption = dataUpdate.Caption
		data.Images = dataUpdate.Images

		var posts Post
		posts.ID = dataUpdate.ID
		txUpdateId := pd.db.Model(&posts).Updates(data)
		if txUpdateId.Error != nil {
			return -1, txUpdateId.Error
		}
		var err error
		return int(txUpdateId.RowsAffected), err
	} else {
		return -1, errors.New("not have access")
	}
}
