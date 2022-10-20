package domain

import (
	"time"
)

type Comments struct {
	ID       uint
	User_ID  uint  
	Post_ID  uint  
	Username string
	Comment  string
	Created_At time.Time 
}

type DataInterface interface {
	AddComment(data Comments) (Comments, error)
	DeleteComent(param, token int) (int, error)
}

type ServiceInterface interface {
	Insert(data Comments) (Comments, error)
	DeleteId(param, token int) (int, error)
}
