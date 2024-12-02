package notification

type NotificationQueryRepository interface {
	Query()
}

type NotificationCommandRepository interface {
	Command()
}
