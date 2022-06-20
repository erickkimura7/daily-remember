package notificationevent

type NotificationService interface {
	AddNotification(model *Event) error
	ListAllNotifications() ([]*Event, error)
}
