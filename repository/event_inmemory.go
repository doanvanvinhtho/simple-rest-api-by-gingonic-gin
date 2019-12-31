package repository

import (
	"errors"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
)

// New an event use in-memory
func New() *EventInMemory {
	return &EventInMemory{
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

// EventInMemory is instance of event repository
type EventInMemory struct {
	events []model.Event
}

// GetOneEvent helps to get a event from in-memory repository
func (e *EventInMemory) GetOneEvent(id string) (*model.Event, error) {
	for _, singleEvent := range e.events {
		if singleEvent.ID == id {
			return &singleEvent, nil
		}
	}

	return nil, errors.New("Event " + id + " not found")
}
