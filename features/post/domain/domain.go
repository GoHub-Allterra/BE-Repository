package domain

import (
	"time"
)

type Post struct {
	// gorm.Model
	ID         uint
	User_ID    uint      `json:"user_id" form:"user_id"`
	Caption    string    `json:"caption" form:"caption"`
	Images     string    `json:"images" form:"images"`
	Created_At time.Time `json:"created_at" form:"created_at"`
	Updated_At time.Time `json:"updated_at" form:"updated_at"`
	Comment    []CommentsCore
	// Deleted_at time.Time `json:"deleted_at" form:"deleted_at"`
	// Post_images []string `json:"post_images" form:"post_images"`
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

type CommentsCore struct {
	ID         uint
	User_ID    uint
	Post_ID    uint
	Comment    string `json:"comment" form:"comment"`
	Created_At time.Time
}

type PostUsecase interface {
	GetAllPosts() (data []Post, err error)
	AddPost(data Post, token int) (int, error)
	SelectById(param int) (data Post, err error)
	UpdatePost(param, token int, data Post) (int, error)
	DeletedPost(param, token int) (int, error)
	GetMyPosts(id uint) ([]Post, error)
}

type PostData interface {
	GetAll() (data []Post, err error)
	Insert(data Post, token int) (int, error)
	GetById(param int) (data Post, err error)
	PutPost(param, token int, data Post) (int, error)
	DeletedId(param, token int) (int, error)
	GetAllPostsByID(id uint) ([]Post, error)
}
