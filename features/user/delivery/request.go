package delivery

import (
	"gohub/features/user/domain"
)

type RegisterFormat struct {
	Nama     string `json:"nama" form:"nama"`
	HP       string `json:"hp" form:"hp"`
	Password string `json:"password" form:"password"`
}

type GetId struct {
	id uint `param:"id"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Nama: cnv.Nama, HP: cnv.HP, Password: cnv.Password}
	case GetId:
		cnv := i.(GetId)
		return domain.Core{ID: cnv.id}
	}
	return domain.Core{}
}
