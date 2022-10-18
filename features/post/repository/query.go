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
