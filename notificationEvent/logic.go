package notificationevent

import (
	"errors"
	"time"

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

	n.notificationRepo.AddEvent(&Event{
		ID:          uuid.String(),
		Title:       "Teste Title",
		Description: "Teste Description",
		DateTime:    time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	return nil
}

func (n *notificationService) ListAllNotifications() ([]*Event, error) {
	allEvents, err := n.notificationRepo.FindAllEvents()

	if err != nil {
		return nil, err
	}

	return allEvents, nil
}
