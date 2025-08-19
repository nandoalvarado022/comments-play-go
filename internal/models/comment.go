package models

import "time"

type Comment struct {
	CommentId      string    `json:"id"`
	UserId         string    `json:"user_id"`
	CommentContend string    `json:"post_content"`
	CreadoEn       time.Time `json:"creado_en"`
}
