package delivery

import "gohub/features/post/domain"

type Request struct {
	Caption string `json:"caption" form:"caption"`
	Images  string `json:"images" form:"images"`
}

func ToDomain(i interface{}) domain.Post {
	switch i.(type) {
	case Request:
		cnv := i.(Request)
		return domain.Post{Caption: cnv.Caption, Images: cnv.Images}
	}
	return domain.Post{}
}
