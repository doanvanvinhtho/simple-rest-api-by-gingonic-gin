package mysql

import (
	"errors"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository"
)

// New an event use MySQL
func New() repository.Event {
	return nil
}

// eventMySQL is instance of event repository
type eventMySQL struct {
}

// GetOneEvent helps to get a event from MySQL repository
func (e *eventMySQL) GetOneEvent(id string) (*model.Event, error) {
	return nil, errors.New("Event " + id + " not found")
}
