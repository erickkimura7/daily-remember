package mockrepository

import (
	"errors"

	nt "github.com/erickkimura7/daily-remember/notificationEvent"
)

type mockRepository struct {
	events []*nt.Event
}

func NewMockRepository() (nt.NotificationRepository, error) {

	repo := &mockRepository{
		events: make([]*nt.Event, 0),
	}

	return repo, nil
}

func (m *mockRepository) AddEvent(event *nt.Event) error {
	m.events = append(m.events, event)

	return nil
}

func (m *mockRepository) FindEventById(id string) (*nt.Event, error) {
	for _, each := range m.events {
		if string(each.ID) == id {
			return each, nil
		}
	}

	return nil, nil
}

func (m *mockRepository) FindAllEvents() ([]*nt.Event, error) {
	return m.events, nil
}

func (m *mockRepository) RemoveEventById(id string) error {
	for index, each := range m.events {
		if string(each.ID) == id {
			m.events = append(m.events[:index], m.events[index+1:]...)
			return nil
		}
	}

	return errors.New("id nao encontrado para ser removido")
}
