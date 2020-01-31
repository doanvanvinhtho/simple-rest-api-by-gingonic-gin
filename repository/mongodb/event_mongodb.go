package mongodb

import (
	"errors"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository"
)

// New an event use MongoDB
func New() repository.Event {
	return nil
}

// eventMongoDB is instance of event repository
type eventMongoDB struct {
}

// GetOneEvent helps to get a event from MongoDB repository
func (e *eventMongoDB) GetOneEvent(id string) (*model.Event, error) {
	return nil, errors.New("Event " + id + " not found")
}
