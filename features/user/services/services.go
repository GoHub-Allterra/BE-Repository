package services

import (
	"errors"
	"gohub/features/user/domain"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &userService{
		qry: repo,
	}
}

func GenerateToken(id uint) string {
	claim := &jwt.MapClaims{
		"authorized": true,
		"id":         id,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	str, err := token.SignedString([]byte("k0D3jW7"))
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	return str
}

func (us *userService) Login(input domain.Core) (domain.Core, string, error) {
	res, err := us.qry.Login(input)
	if err != nil {
		log.Error(err.Error(), "username not found")
		return domain.Core{}, "", err
	}

	pass := domain.Core{Password: res.Password}
	check := bcrypt.CompareHashAndPassword([]byte(pass.Password), []byte(input.Password))
	if check != nil {
		log.Error(check, "wrong password")
		return domain.Core{}, "",check
	}
	token := GenerateToken(res.ID)

	return res, token, err
}

func (us *userService) UpdateUser(input domain.Core)(domain.Core, error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)	
	if err != nil {
		log.Error(err.Error())
		return domain.Core{}, errors.New("cannot encrypt password")
	}

	input.Password = string(generate)
	res, err := us.qry.Edit(input)
		if err != nil {
			return domain.Core{}, err
		}
	return res, nil
}

func (us *userService) DeleteUser(id uint)(domain.Core, error) {
	res, err := us.qry.Delete(id)
	if err != nil {
		return domain.Core{}, err
	}
	return res, err
}


func (us *userService) AddUser(newUser domain.Core) (domain.Core, error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Error(err.Error())
		return domain.Core{}, errors.New("cannot encrypt password")
	}

	newUser.Password = string(generate)
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
		return domain.Core{}, errors.New("no data")
	}

	return res, nil
}

// func (us *userService) ShowAllUser() ([]domain.Core, error) {
// 	res, err := us.qry.GetAll()
// 	if err != nil {
// 		log.Error(err.Error())
// 		if strings.Contains(err.Error(), "table") {
// 			return nil, errors.New("database error")
// 		} else if strings.Contains(err.Error(), "found") {
// 			return nil, errors.New("no data")
// 		}
// 	}

// 	if len(res) == 0 {
// 		log.Info("no data")
// 		return nil, errors.New("no data")
// 	}

// 	return res, nil
// }
