package notificationevent

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNotificationNotFound = errors.New("Notification Not Found")
	ErrNotificationInvalid  = errors.New("Notification Invalid")
)

type notificationService struct {
	notificationRepo NotificationRepository
}

func NewNotificationService(notificationRepo NotificationRepository) NotificationService {
	return &notificationService{
		notificationRepo,
	}
}

func (n *notificationService) AddNotification(model *Event) error {

	uuid, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	model.ID = uuid.String()

	err = n.notificationRepo.AddEvent(model)

	return err
}

func (n *notificationService) ListAllNotifications() ([]*Event, error) {
	allEvents, err := n.notificationRepo.FindAllEvents()

	if err != nil {
		return nil, err
	}

	return allEvents, nil
}

func (n *notificationService) RemoveNotification(id string) error {
	err := n.notificationRepo.RemoveEventById(id)

	return err
}
