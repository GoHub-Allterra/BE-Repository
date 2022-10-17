package services

import (
	"errors"
	"gohub/features/user/domain"
	"strings"

	"github.com/labstack/gommon/log"
)

type userService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &userService{
		qry: repo,
	}
}

func (us *userService)UpdateUser(id uint, input domain.Core)(domain.Core, error) {
	res, err := us.qry.Edit(id, input)
	if err != nil {
		log.Error(err.Error())
		return domain.Core{}, err
	}
	return res, nil
}

func (us *userService)DeleteUser(id uint)(domain.Core, error) {
	res, err := us.qry.Delete(id)
	if err != nil {
		return domain.Core{}, err
	}
	return res, err
}

// func GenerateToken(id uint) (string, error) {
// 	claim := &jwt.MapClaims{
// 		"authorized": true,
// 		"id":         id,
// 		"exp":        time.Now().Add(time.Hour * 1).Unix(),
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

// 	str, err := token.SignedString([]byte("rahasia"))
// 	if err != nil {
// 		log.Error(err.Error())
// 		return "", errors.New("failed build token")
// 	}
// 	return str, nil
// }

func (us *userService) AddUser(newUser domain.Core) (domain.Core, error) {
	// generate, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	// if err != nil {
	// 	log.Error(err.Error())
	// 	return domain.Core{}, errors.New("cannot encript password")
	// }

	// newUser.Password = string(generate)
	res, err := us.qry.Insert(newUser)

	if err != nil {
		return domain.Core{}, errors.New("some problem on database")
	}

	return res, nil
}

func (us *userService) Get(ID uint) (domain.Core, error) {
	res, err := us.qry.Get(ID)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}

	return res, nil
}

func (us *userService) ShowAllUser() ([]domain.Core, error) {
	res, err := us.qry.GetAll()
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return nil, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, errors.New("no data")
		}
	}

	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New("no data")
	}

	return res, nil
}
