package service

import (
	"net/http"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository"
)

// New an event in-memory service
func New(r *repository.EventInMemory) *EventInmemory {
	return &EventInmemory{
		repo: r,
	}
}

// EventInmemory is instance of event service
type EventInmemory struct {
	repo *repository.EventInMemory
}

// GetOneEvent helps to get an event
func (s *EventInmemory) GetOneEvent(id string) model.Response {
	var response model.Response

	if event, err := s.repo.GetOneEvent(id); err != nil {
		response = model.Response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		}
	} else {
		response = model.Response{
			Code:    http.StatusOK,
			Message: "",
			Data:    *event,
		}
	}

	return response
}

// GetAllEvent helps to get all events
func (s *EventInmemory) GetAllEvent() []model.Response {
	return nil
}
