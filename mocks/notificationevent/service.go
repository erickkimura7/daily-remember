// Code generated by MockGen. DO NOT EDIT.
// Source: notificationEvent/service.go

// Package mock_notificationevent is a generated GoMock package.
package mock_notificationevent

import (
	reflect "reflect"

	notificationevent "github.com/erickkimura7/daily-remember/notificationEvent"
	gomock "github.com/golang/mock/gomock"
)

// MockNotificationService is a mock of NotificationService interface.
type MockNotificationService struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationServiceMockRecorder
}

// MockNotificationServiceMockRecorder is the mock recorder for MockNotificationService.
type MockNotificationServiceMockRecorder struct {
	mock *MockNotificationService
}

// NewMockNotificationService creates a new mock instance.
func NewMockNotificationService(ctrl *gomock.Controller) *MockNotificationService {
	mock := &MockNotificationService{ctrl: ctrl}
	mock.recorder = &MockNotificationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationService) EXPECT() *MockNotificationServiceMockRecorder {
	return m.recorder
}

// AddNotification mocks base method.
func (m *MockNotificationService) AddNotification(model *notificationevent.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNotification", model)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNotification indicates an expected call of AddNotification.
func (mr *MockNotificationServiceMockRecorder) AddNotification(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNotification", reflect.TypeOf((*MockNotificationService)(nil).AddNotification), model)
}

// ListAllNotifications mocks base method.
func (m *MockNotificationService) ListAllNotifications() ([]*notificationevent.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllNotifications")
	ret0, _ := ret[0].([]*notificationevent.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllNotifications indicates an expected call of ListAllNotifications.
func (mr *MockNotificationServiceMockRecorder) ListAllNotifications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllNotifications", reflect.TypeOf((*MockNotificationService)(nil).ListAllNotifications))
}

// RemoveNotification mocks base method.
func (m *MockNotificationService) RemoveNotification(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveNotification", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNotification indicates an expected call of RemoveNotification.
func (mr *MockNotificationServiceMockRecorder) RemoveNotification(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNotification", reflect.TypeOf((*MockNotificationService)(nil).RemoveNotification), id)
}
