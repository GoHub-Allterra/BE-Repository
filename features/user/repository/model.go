package repository

import (
	"gohub/features/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string
	Username        string
	Email           string
	Password        string
	Profile_picture string
	HP              string
	Bio             string
}

func FromDomain(du domain.Core) User {
	return User{
		Model:           gorm.Model{ID: du.ID},
		Name:            du.Name,
		Username:        du.Username,
		Password:        du.Password,
		Profile_picture: du.Profile_picture,
		HP:              du.Profile_picture,
		Bio:             du.Bio,
	}
}

func ToDomain(u User) domain.Core {
	return domain.Core{
		ID:              u.ID,
		Name:            u.Name,
		Username:        u.Username,
		Password:        u.Password,
		Profile_picture: u.Profile_picture,
		HP:              u.HP,
		Bio:             u.Bio,
	}
}

// func ToDomainArray(au []User) []domain.Core {
// 	var res []domain.Core
// 	for _, val := range au {
// 		res = append(res, domain.Core{ID: val.ID, Name: val.Name, Username: val.Name, Password: val.Password, Profile_picture: val.Profile_picture})
// 	}

// 	return res
// }
