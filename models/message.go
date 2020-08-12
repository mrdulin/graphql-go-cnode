package models

type MessageId struct {
	ID string `json:"id" mapstructure:"id"`
}

type Message struct {
	MessageId `mapstructure:",squash"`
	Type      string    `json:"type"`
	HasRead   bool      `json:"has_read" mapstructure:"has_read"`
	Author    User      `json:"author"`
	Topic     TopicBase `json:"topic"`
	Reply     ReplyBase `json:"reply"`
	CreateAt  string    `json:"create_at" mapstructure:"create_at"`
}

type GetMessagesResponse struct {
	HasReadMessages    []Message `json:"has_read_messages" mapstructure:"has_read_messages"`
	HasnotReadMessages []Message `json:"hasnot_read_messages" mapstructure:"hasnot_read_messages"`
}

type MarkAllRequest struct {
	AccessToken string `json:"accesstoken"`
}

type MarkAllMessagesResponse struct {
	MarkedMsgs []MessageId `json:"marked_msgs"`
}
