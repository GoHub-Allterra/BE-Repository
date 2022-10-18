package delivery

import (
	"gohub/features/user/domain"
)

type RegisterFormat struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
}

type UpdateFormat struct {
	ID       uint
	Name     string `form:"name" json:"name"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Email    string `form:"email" json:"email"`
	Images   string `form:"images" json:"images"`
	HP       string `form:"hp" json:"hp"`
	Bio      string `form:"bio" json:"bio"`
}

type LoginFormat struct {
	Usename  string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
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
		return domain.Core{ID: cnv.ID,Name: cnv.Name, Username: cnv.Username, Password: cnv.Password, Images: cnv.Images, Email: cnv.Email, HP: cnv.HP, Bio: cnv.Bio}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.Core{Username: cnv.Usename, Password: cnv.Password}
	}
	return domain.Core{}
}
