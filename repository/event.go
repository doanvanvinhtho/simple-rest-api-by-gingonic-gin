package repository

import "github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"

// Event is abstract of event repository
type Event interface {
	GetOneEvent(id string) (*model.Event, error)
}
