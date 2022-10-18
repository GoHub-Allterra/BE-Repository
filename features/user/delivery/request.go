package delivery

import (
	"gohub/features/user/domain"
)

type RegisterFormat struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UpdateFormat struct {
	ID       uint
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	HP       string `json:"hp"`
	Bio      string `json:"bio"`
}

type LoginFormat struct {
	Usename  string `json:"username"`
	Password string `json:"password"`
}

type AddPhotosFormat struct {
	ID     uint   `form:"id"`
	Images string `form:"images"`
}

type GetId struct {
	id uint `param:"id"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Name: cnv.Name, Username: cnv.Username, Password: cnv.Password, Email: cnv.Email}
	case GetId:
		cnv := i.(GetId)
		return domain.Core{ID: cnv.id}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{Name: cnv.Name, Username: cnv.Username, Password: cnv.Password, Email: cnv.Email, HP: cnv.HP, Bio: cnv.Bio}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.Core{Username: cnv.Usename, Password: cnv.Password}
	}
	return domain.Core{}
}
