package repository

import (
	"errors"
	"gohub/features/comments/domain"

	"github.com/labstack/gommon/log"
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

func (cd *commentData) AddComment(data domain.Comments) (domain.Comments, error) {
	var input Comments
	input = ToEntity(data)

	res := cd.db.Create(&input)
	if res.Error != nil {
		log.Error("ERROR QUERY")
		return domain.Comments{}, res.Error
	}

	return data, res.Error
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
