package delivery

import "time"

type GetAllPost struct {
	ID       uint
	User_ID  uint `json:"user_id" form:"user_id"`
	Username string
	Images   string
	Content  string
	// Profile_picture_path string
	// Caption    string
	// Created_At time.Time `json:"created_at" form:"created_at"`
	// Updated_At time.Time `json:"updated_at" form:"updated_at"`
	// Post_Images          []string
}

type GetComments struct {
	ID                   uint
	Username             string
	Profile_picture_path string
	Caption              string
	Created_At           time.Time
}

type GetSpecificPost struct {
	ID                   uint
	User_ID              uint
	Username             string
	Profile_picture_path string
	Caption              string
	Created_At           time.Time
	Updated_At           time.Time
	Post_Images          []string
	Comments             []GetComments
}
