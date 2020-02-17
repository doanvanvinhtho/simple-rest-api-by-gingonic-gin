package repository

import "github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"

// Event is abstract of event repository
type Event interface {
	GetOneEvent(id string) (ev *model.Event, err error)
	GetAllEvent() (ev []*model.Event, err error)
	AddEvent(ev *model.Event) (id string, err error)
	UpdateEvent(ev *model.Event) (id string, err error)
	DeleteEvent(id string) error
}
