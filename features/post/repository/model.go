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
		ID:         p.ID,
		User_ID:    p.User_ID,
		Caption:    p.Caption,
		Images:     p.Images,
		Created_At: p.CreatedAt,
		Updated_At: p.UpdatedAt,
	}
}

func ToEntity(data domain.Post) Post {
	return Post{
		User_ID: data.User_ID,
		Caption: data.Caption,
		Images:  data.Images,
		// CreatedAt: data.Created_At,
		// UpdatedAt: data.Updated_At,
	}
}

func (dataPost *Post) toPostUser() domain.Post {

	dataPostCore := domain.Post{
		ID:         dataPost.ID,
		User_ID:    dataPost.User_ID,
		Images:     dataPost.Images,
		Caption:    dataPost.Caption,
		Created_At: dataPost.CreatedAt,
		Updated_At: dataPost.UpdatedAt,
	}

	return dataPostCore

}

func toPostList(data []Post) []domain.Post {
	var dataCore []domain.Post
	for i := 0; i < len(data); i++ {
		dataCore = append(dataCore, data[i].toPostUser())
	}
	return dataCore
}
