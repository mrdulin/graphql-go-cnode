package models

type TopicBase struct {
	ID          string `json:"id"`
	AuthorID    string `json:"author_id"`
	Tab         string `json:"tab"`
	Content     string `json:"content"`
	Title       string `json:"title"`
	LastReplyAt string `json:"last_reply_at"`
	Good        bool   `json:"good"`
	Top         bool   `json:"top"`
	ReplyCount  int    `json:"reply_count"`
	VisitCount  int    `json:"visit_count"`
	CreateAt    string `json:"create_at"`
}

type Topic struct {
	TopicBase
	Author User `json:"author"`
}

type TopicDetail struct {
	Topic
	Replies []Reply `json:"replies"`
}

type RecentTopic struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	LastReplyAt string `json:"last_reply_at"`
	Author      User   `json:"author"`
}