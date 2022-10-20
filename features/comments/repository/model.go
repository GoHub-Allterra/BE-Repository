package repository

import (
	"gohub/features/comments/domain"

	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	User_ID  uint   `json:"user_id" form:"user_id"`
	Post_ID  uint   `json:"post_id" form:"post_id"`
	Username string `json:"username" form:"username"`
	Comment  string `json:"comment" form:"comment"`
	// Created_At time.Time `json:"created_at" form:"created_at"`
	// Updated_At time.Time `json:"updated_at" form:"updated_at"`
}

func (cm *Comments) ToDomain() domain.Comments {
	return domain.Comments{
		ID:         cm.ID,
		User_ID:    cm.User_ID,
		Post_ID:    cm.Post_ID,
		Username:   cm.Username,
		Comment:    cm.Comment,
		Created_At: cm.CreatedAt,
	}
}

func ToEntity(data domain.Comments) Comments {
	return Comments{
		User_ID:  data.User_ID,
		Post_ID:  data.Post_ID,
		Username: data.Username,
		Comment:  data.Comment,
		// CreatedAt: data.Created_At,
		// UpdatedAt: data.Updated_At,
	}
}

func (dataPost *Comments) toPostUser() domain.Comments {

	dataPostCore := domain.Comments{
		ID:         dataPost.ID,
		User_ID:    dataPost.User_ID,
		Post_ID:    dataPost.Post_ID,
		Username:   dataPost.Username,
		Comment:    dataPost.Comment,
		Created_At: dataPost.CreatedAt,
	}

	return dataPostCore

}

func toPostList(data []Comments) []domain.Comments {
	var dataCore []domain.Comments
	for i := 0; i < len(data); i++ {
		dataCore = append(dataCore, data[i].toPostUser())
	}
	return dataCore
}
