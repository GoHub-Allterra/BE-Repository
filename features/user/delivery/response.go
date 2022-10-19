package delivery

import "gohub/features/user/domain"

func SuccessResponseWithData(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func SuccessResponseNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func FailResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

type GetResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
	HP       string `json:"hp"`
	Bio      string `json:"bio"`
	Images   string `json:"images"`
}

type LoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "get":
		cnv := core.(domain.Core)
		res = GetResponse{ID: cnv.ID, Name: cnv.Name, Username: cnv.Username, Email: cnv.Name, HP: cnv.HP, Bio: cnv.Bio, Images: cnv.Images}
	case "login":
		cnv := core.(domain.Core)
		res = LoginResponse{ID: cnv.ID, Name: cnv.Name, Token: cnv.Token}
	}
	return res
}
