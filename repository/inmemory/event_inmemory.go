package inmemory

import (
	"errors"
	"sync"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository"
)

// New an event use in-memory
func New() repository.Event {
	return &eventInMemory{
		events: []model.Event{
			{
				ID:          "id_go",
				Title:       "Go",
				Description: "https://golang.org/",
			},
			{
				ID:          "id_gin",
				Title:       "Gin",
				Description: "https://github.com/gin-gonic/gin",
			},
		},
	}
}

// eventInMemory is instance of event repository
type eventInMemory struct {
	events []model.Event
	mu     sync.Mutex
}

// GetOneEvent helps to get an event from repository
func (e *eventInMemory) GetOneEvent(id string) (*model.Event, error) {
	for _, singleEvent := range e.events {
		if singleEvent.ID == id {
			return &singleEvent, nil
		}
	}

	return nil, errors.New("Event " + id + " not found")
}

// GetAllEvent helps to get all events from repository
func (e *eventInMemory) GetAllEvent() ([]model.Event, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	result := make([]model.Event, len(e.events))
	for i, singleEvent := range e.events {
		result[i] = singleEvent
	}

	return result, nil
}

// AddEvent helps to add new event into repository
func (e *eventInMemory) AddEvent(ev *model.Event) (string, error) {
	if ev == nil {
		return "", errors.New("The event need to add is nil")
	}

	e.events = append(e.events, *ev)

	return ev.ID, nil
}

// UpdateEvent helps to update an event in repository
func (e *eventInMemory) UpdateEvent(ev *model.Event) (string, error) {
	if ev == nil {
		return "", errors.New("The event need to update is nil")
	}

	for i := 0; i < len(e.events); i++ {
		if e.events[i].ID == ev.ID {
			e.events[i].Title = ev.Title
			e.events[i].Description = ev.Description

			return ev.ID, nil
		}
	}

	return "", errors.New("Event " + ev.ID + " not found")
}

// DeleteEvent helps to delete an event in repository
func (e *eventInMemory) DeleteEvent(id string) error {
	if len(id) <= 0 {
		return errors.New("The event id is empty")
	}

	for i := 0; i < len(e.events); i++ {
		if e.events[i].ID == id {
			e.events = append(e.events[0:i], e.events[i+1:]...)
			return nil
		}
	}

	return errors.New("Event " + id + " not found")
}
