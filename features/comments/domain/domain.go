package domain

import (
	"time"
)

type Comments struct {
	ID       uint
	User_ID  uint   `json:"user_id" form:"user_id"`
	Post_ID  uint   `json:"post_id" form:"post_id"`
	Username string `json:"username" form:"username"`
	Comment  string `json:"comment" form:"comment"`
	// Images     string    `json:"images" form:"images"`
	Created_At time.Time `json:"created_at" form:"created_at"`
	// Post       Post
}

type Post struct {
	ID         uint      `json:"id" form:"id"`
	User_Id    string    `json:"user_id" form:"user_id"`
	Images     string    `json:"images" form:"images"`
	Content    string    `json:"content" form:"content"`
	Created_At time.Time `json:"created_at" form:"created_at"`
	// User       User
	Comment []Comments
}

type User struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Hp       string `json:"hp" form:"hp"`
	Bio      string `json:"bio" form:"bio"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Post     []Post
}

type DataInterface interface {
	AddComment(data Comments) (Comments, error)
	DeleteComent(param, token int) (int, error)
}

type ServiceInterface interface {
	Insert(data Comments) (Comments, error)
	DeleteId(param, token int) (int, error)
}
