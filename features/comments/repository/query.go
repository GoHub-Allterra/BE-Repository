package repository

import (
	"errors"
	"gohub/features/comments/domain"

	"gorm.io/gorm"
)

type commentData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.DataInterface {
	return &commentData{
		db: DB,
	}
}

func (cd *commentData) AddComment(data domain.Comments, param int) (domain.Comments, error) {
	var datacheck domain.Comments
	txcheck := cd.db.Where("ID=?", param).First(&datacheck)
	if txcheck.Error != nil {
		return domain.Comments{}, errors.New("error tx")
	}

	if int(datacheck.ID) != param {
		return datacheck, nil
	}

	dataModel := ToEntity(data)
	dataModel.User_ID = uint(param)
	tx := cd.db.Create(&dataModel)
	if tx.Error != nil {
		return domain.Comments{}, tx.Error
	}

	return datacheck, nil

	//===============================
	// var datacheck domain.Comments
	// txcheck := cd.db.Where("ID=?", param).First(&datacheck)
	// if txcheck.Error != nil {
	// 	return domain.Comments{}, errors.New("error tx")
	// }

	// if int(datacheck.ID) != param {
	// 	return datacheck, nil
	// }

	// var commentData Comments = ToEntity(data)
	// err := cd.db.Create(&commentData).Error
	// if err != nil {
	// 	return domain.Comments{}, err
	// }

	// if int(commentData.ID) != param {
	// 	return domain.Comments{}, nil
	// }
	// // return datacheck, nil
	// return commentData.ToDomain(), nil
}
func (cd *commentData) DeleteComent(param, token int) (int, error) {

	dataCart := Comments{}
	idCheck := cd.db.First(&dataCart, param)
	if idCheck.Error != nil {
		return 0, idCheck.Error
	}
	if uint(token) != dataCart.User_ID {
		return -1, errors.New("you don't have access")
	}
	result := cd.db.Delete(&Comments{}, param)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to delete data")
	}
	return int(result.RowsAffected), nil

}
