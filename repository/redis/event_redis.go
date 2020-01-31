package redis

import (
	"errors"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository"
)

// New an event use redis
func New() repository.Event {
	return nil
}

// eventRedis is instance of event repository
type eventRedis struct {
}

// GetOneEvent helps to get a event from redis repository
func (e *eventRedis) GetOneEvent(id string) (*model.Event, error) {
	return nil, errors.New("Event " + id + " not found")
}
