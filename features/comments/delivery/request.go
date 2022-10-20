package delivery

import "gohub/features/comments/domain"

type CommentFormat struct {
	IdPost uint
	IdUser uint
	Comment string `json:"comment" form:"comment"`

}

func ToDomain(i CommentFormat) domain.Comments {
	// var cnv CommentFormat
	return domain.Comments{Comment: i.Comment, User_ID: i.IdUser, Post_ID: i.IdPost}
}
