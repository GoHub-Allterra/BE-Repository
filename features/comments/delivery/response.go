package delivery

import "time"

type GetComments struct {
	ID         uint
	Username   string
	Profile    string
	Caption    string
	Created_At time.Time
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
