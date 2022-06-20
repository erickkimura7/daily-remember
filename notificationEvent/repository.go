package notificationevent

type NotificationRepository interface {
	AddEvent(events *Event) error
	FindEventById(id string) (*Event, error)
	FindAllEvents() ([]*Event, error)
}
