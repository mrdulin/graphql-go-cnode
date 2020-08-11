package models

type TopicTabEnum string

const (
	TOPIC_TAB_ASK   TopicTabEnum = "ask"
	TOPIC_TAB_JOB   TopicTabEnum = "job"
	TOPIC_TAB_GOOD  TopicTabEnum = "good"
	TOPIC_TAB_SHARE TopicTabEnum = "share"
)

type TopicBase struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	LastReplyAt string `json:"last_reply_at"`
}

type Topic struct {
	TopicBase
	AuthorID   string `json:"author_id"`
	Tab        string `json:"tab"`
	Content    string `json:"content"`
	Good       bool   `json:"good"`
	Top        bool   `json:"top"`
	ReplyCount int    `json:"reply_count"`
	VisitCount int    `json:"visit_count"`
	CreateAt   string `json:"create_at"`
	Author     User   `json:"author"`
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
