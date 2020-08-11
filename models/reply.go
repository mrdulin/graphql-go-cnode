package models

import "time"

type ReplyBase struct {
	ID       string    `json:"id"`
	Content  string    `json:"content"`
	Ups      []string  `json:"ups"`
	CreateAt time.Time `json:"create_at"`
}

type Reply struct {
	Author  User        `json:"author"`
	ReplyID interface{} `json:"reply_id"`
	IsUped  bool        `json:"is_uped"`
}
