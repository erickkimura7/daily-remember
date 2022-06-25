package notificationevent

type NotificationService interface {
	AddNotification(model *Event) error
	RemoveNotification(id string) error
	ListAllNotifications() ([]*Event, error)
}
