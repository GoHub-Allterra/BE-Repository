package repository

import (
	"gohub/features/post/domain"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	User_ID uint   `json:"user_id" form:"user_id"`
	Images  string `json:"images" form:"images"`
	Caption string `json:"caption" form:"caption"`
	// Created_At time.Time `json:"created_at" form:"created_at"`
	// Updated_At time.Time `json:"updated_at" form:"updated_at"`
}

func (p *Post) ToDomain() domain.Post {
	return domain.Post{
		ID:      p.ID,
		User_ID: p.User_ID,
		Caption: p.Caption,
		Images:  p.Images,
		// Created_At: p.Created_At,
		// Updated_At: p.Updated_At,
	}
}

func ToEntity(data domain.Post) Post {
	return Post{
		User_ID: data.User_ID,
		Caption: data.Caption,
		Images:  data.Images,
		// Created_At: data.Created_At,
		// Updated_At: data.Updated_At,
	}
}

func (dataPost *Post) toPostUser() domain.Post {

	dataBookCore := domain.Post{
		ID:      dataPost.ID,
		User_ID: dataPost.User_ID,
		Images:  dataPost.Images,
		Caption: dataPost.Caption,
		// Created_At: dataPost.Created_At,
		// Updated_At: dataPost.Updated_At,
	}

	return dataBookCore

}

func toPostList(data []Post) []domain.Post {
	var dataCore []domain.Post
	for i := 0; i < len(data); i++ {
		dataCore = append(dataCore, data[i].toPostUser())
	}
	return dataCore
}
