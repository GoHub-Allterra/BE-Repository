package delivery

import "gohub/features/comments/domain"

type Request struct {
	Comment string `json:"comment"`
}

func ToDomain(i interface{}) domain.Comments {
	switch i.(type) {
	case Request:
		cnv := i.(Request)
		return domain.Comments{Comment: cnv.Comment}
	}
	return domain.Comments{}
}
