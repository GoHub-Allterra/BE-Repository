package domain

import (
	"time"
)

type Post struct {
	ID         uint
	User_ID    uint   `json:"user_id" form:"user_id"`
	Caption    string `json:"caption" form:"caption"`
	Images     string `json:"images" form:"images"`
	Created_At time.Time
	Updated_At time.Time
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
}

type Comment struct {
	ID         uint
	User_ID    uint
	Post_ID    uint
	Caption    string `json:"caption" form:"caption"`
	Created_At time.Time
}

type PostUsecase interface {
	GetAllPosts() (data []Post, err error)
	AddPost(data Post, token int) (int, error)
	// AddPostImages(post []string, postID uint) error
	// GetMyPosts(id uint) ([]Post, User, [][]string, error)
	// GetSpecificPost(id uint) (Post, User, []string, []Comment, []User, error)
	// UpdatePost(id uint, updateData Post) (Post, error)
	// DeletePost(id uint, userID uint) error
}

type PostData interface {
	GetAll() (data []Post, err error)
	Insert(data Post, token int) (int, error)
	// InsertPostImages(post []string, postID uint) error
	// GetAllPostsByID(id uint) ([]Post, User, [][]string, error)
	// GetPostByID(id uint) (Post, User, []string, []Comment, []User, error)
	// Update(id uint, updatePost Post) (Post, error)
	// Delete(id uint, userID uint) error
}
