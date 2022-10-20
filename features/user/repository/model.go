package repository

import (
	"gohub/features/user/domain"
	user "gohub/features/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Username string
	Email    string
	Password string
	Images   string
	HP       string
	Bio      string
	Post     []Post
}

type Post struct {
	gorm.Model
	UserID  uint
	Images  string
	Content string
}

func (dataPostUser *Post) toPostUser() user.PostCore {

	dataUser := user.PostCore{
		ID:      dataPostUser.ID,
		UserID:  dataPostUser.UserID,
		Images:  dataPostUser.Images,
		Content: dataPostUser.Content,
	}

	return dataUser

}

func FromDomain(du domain.Core) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Name:     du.Name,
		Username: du.Username,
		Email:    du.Email,
		Password: du.Password,
		Images:   du.Images,
		HP:       du.HP,
		Bio:      du.Bio,
	}
}

func ToDomain(u User) domain.Core {
	return domain.Core{
		ID:       u.ID,
		Name:     u.Name,
		Username: u.Username,
		Password: u.Password,
		Images:   u.Images,
		HP:       u.HP,
		Bio:      u.Bio,
	}
}
