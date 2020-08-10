package models

import "time"

type Reply struct {
	ID       string      `json:"id"`
	Author   User        `json:"author"`
	Content  string      `json:"content"`
	Ups      []string    `json:"ups"`
	CreateAt time.Time   `json:"create_at"`
	ReplyID  interface{} `json:"reply_id"`
	IsUped   bool        `json:"is_uped"`
}
