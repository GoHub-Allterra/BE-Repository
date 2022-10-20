package repository

import (
	"gohub/features/user/domain"
	"log"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) AddPhotos(input domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(input)
	if err := rq.db.Model(&cnv).Where("id = ?", input.ID).Update("images", input.Images).Error; err != nil {
		log.Fatal("error update data")
		return domain.Core{}, err
	}
	input = ToDomain(cnv)
	return input, nil
}

func (rq *repoQuery) Login(input domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(input)
	if err := rq.db.Where("username = ?", cnv.Username).Find(&cnv).Error; err != nil {
		log.Fatal("error get data")
		return domain.Core{}, err
	}
	input = ToDomain(cnv)
	return input, nil
}

func (rq *repoQuery) Delete(id uint) (domain.Core, error) {
	if err := rq.db.Where("id = ?", id).Delete(&User{}); err != nil {
		return domain.Core{}, err.Error
	}
	return domain.Core{}, nil
}

func (rq *repoQuery) Insert(newUser domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(newUser)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}

	// selesai dari DB
	newUser = ToDomain(cnv)
	return newUser, nil
}
func (rq *repoQuery) Edit(input domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(input)
	if err := rq.db.Where("id = ?", 11).Updates(User{Name: input.Name, HP: input.HP, Password: input.Password, Username: input.Username,
		Email: input.Email, Bio: input.Bio}).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	input = ToDomain(cnv)
	return input, nil
}
func (rq *repoQuery) Get(ID uint) (domain.Core, error) {
	var resQry User
	if err := rq.db.First(&resQry, "ID = ?", ID).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	res := ToDomain(resQry)
	return res, nil
}

// func (rq *repoQuery) Login(user domain.Core) (domain.Core, error) {
// 	var dest User
// 	if err := rq.db.First(&dest, "username = ? AND password = ?", user.Username, user.Password).Error ;err != nil {
// 		return domain.Core{}, err
// 	}

// 	res := ToDomain(dest)
// 	return res, nil

// }

// func (rq *repoQuery) GetAll() ([]domain.Core, error) {
// 	var resQry []User
// 	if err := rq.db.Find(&resQry).Error; err != nil {
// 		return nil, err
// 	}
// 	// selesai dari DB
// 	res := ToDomainArray(resQry)
// 	return res, nil
// }
