package service

import (
	"net/http"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository"
)

// New an event service
func New(r repository.Event) Event {
	return &eventImpl{
		repo: r,
	}
}

// EventImpl is instance of event service
type eventImpl struct {
	repo repository.Event
}

// GetOneEvent helps to get an event
func (s *eventImpl) GetOneEvent(id string) *model.Response {
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

	return &response
}

// GetAllEvent helps to get all events
func (s *eventImpl) GetAllEvent() *model.Response {
	var response model.Response

	if events, err := s.repo.GetAllEvent(); err != nil {
		response = model.Response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		}
	} else {
		if events == nil || len(events) <= 0 {
			response = model.Response{
				Code:    http.StatusOK,
				Message: "",
				Data:    nil,
			}
		} else {
			response = model.Response{
				Code:    http.StatusOK,
				Message: "",
				Data:    events,
			}
		}
	}

	return &response
}

// AddEvent helps to add new event
func (s *eventImpl) AddEvent(ev *model.Event) *model.Response {
	var response model.Response

	if _, err := s.repo.AddEvent(ev); err != nil {
		response = model.Response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		}
	} else {
		response = model.Response{
			Code:    http.StatusOK,
			Message: "",
			Data:    nil,
		}
	}

	return &response
}

// UpdateEvent helps to update an event
func (s *eventImpl) UpdateEvent(ev *model.Event) *model.Response {
	var response model.Response

	if _, err := s.repo.UpdateEvent(ev); err != nil {
		response = model.Response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		}
	} else {
		response = model.Response{
			Code:    http.StatusOK,
			Message: "",
			Data:    nil,
		}
	}

	return &response
}

// DeleteEvent helps to delete an event
func (s *eventImpl) DeleteEvent(id string) *model.Response {
	var response model.Response

	if err := s.repo.DeleteEvent(id); err != nil {
		response = model.Response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		}
	} else {
		response = model.Response{
			Code:    http.StatusOK,
			Message: "",
			Data:    nil,
		}
	}

	return &response
}
