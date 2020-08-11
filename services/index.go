package services

type Container struct {
	UserService    UserService
	TopicService   TopicService
	MessageService MessageService
}
