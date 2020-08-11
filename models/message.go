package models

import "github.com/mrdulin/graphql-go-cnode/utils"

type MessageId struct {
	ID string `json:"id"`
}

type Message struct {
	MessageId
	Type     string    `json:"type"`
	HasRead  bool      `json:"has_read"`
	Author   User      `json:"author"`
	Topic    TopicBase `json:"topic"`
	Reply    ReplyBase `json:"reply"`
	CreateAt string    `json:create_at`
}

type Messages struct {
	HasReadMessages    []Message `json:"has_read_messages"`
	HasnotReadMessages []Message `json:"hasnot_read_messages"`
}

type MarkAllMessagesResponse struct {
	utils.ResponseStatus
	MarkedMsgs []MessageId `json:"marked_msgs"`
}
