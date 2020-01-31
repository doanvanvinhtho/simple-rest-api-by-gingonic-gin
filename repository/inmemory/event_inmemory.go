package inmemory

import (
	"errors"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository"
)

// New an event use in-memory
func New() repository.Event {
	return eventInMemory{
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
}

// GetOneEvent helps to get an event from in-memory repository
func (e eventInMemory) GetOneEvent(id string) (*model.Event, error) {
	for _, singleEvent := range e.events {
		if singleEvent.ID == id {
			return &singleEvent, nil
		}
	}

	return nil, errors.New("Event " + id + " not found")
}
