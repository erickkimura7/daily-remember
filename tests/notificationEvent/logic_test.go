package notificationevent

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	mock_notificationevent "github.com/erickkimura7/daily-remember/mocks/notificationevent"
	nt "github.com/erickkimura7/daily-remember/notificationEvent"
)

func TestNotificationService_AddNotification(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_notificationevent.NewMockNotificationRepository(ctrl)

	service := nt.NewNotificationService(mockRepo)

	t.Run("Test success", func(t *testing.T) {
		mockRepo.EXPECT().AddEvent(gomock.Any()).Return(nil)
		err := service.AddNotification(&nt.Event{})

		require.Nil(t, err)
	})

	t.Run("Test error", func(t *testing.T) {
		mockRepo.EXPECT().AddEvent(gomock.Any()).Return(errors.New("test error"))
		err := service.AddNotification(&nt.Event{})

		require.Error(t, err)
	})

}

func TestNotificationService_ListAllNotifications(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_notificationevent.NewMockNotificationRepository(ctrl)

	service := nt.NewNotificationService(mockRepo)

	t.Run("Test success", func(t *testing.T) {
		emptyList := make([]*nt.Event, 0)

		mockRepo.EXPECT().FindAllEvents().Return(emptyList, nil)
		got, err := service.ListAllNotifications()

		require.Nil(t, err)
		require.Equal(t, emptyList, got)
	})

	t.Run("Test error", func(t *testing.T) {

		mockRepo.EXPECT().FindAllEvents().Return(nil, errors.New("test error"))
		got, err := service.ListAllNotifications()

		require.Nil(t, got)
		require.Error(t, err)
	})

}

func TestNotificationService_RemoveNotification(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_notificationevent.NewMockNotificationRepository(ctrl)

	service := nt.NewNotificationService(mockRepo)

	randomId := uuid.New().String()

	t.Run("Test success", func(t *testing.T) {
		mockRepo.EXPECT().RemoveEventById(gomock.Any()).Return(nil)
		err := service.RemoveNotification(randomId)

		require.Nil(t, err)
	})

	t.Run("Test error", func(t *testing.T) {
		mockRepo.EXPECT().RemoveEventById(gomock.Any()).Return(errors.New("teste error"))
		err := service.RemoveNotification(randomId)

		require.Error(t, err)
	})

}
