package service

import "github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"

// Event service interface
type Event interface {
	GetOneEvent(id string) *model.Response
	GetAllEvent() *model.Response
	AddEvent(ev *model.Event) *model.Response
	UpdateEvent(ev *model.Event) *model.Response
	DeleteEvent(id string) *model.Response
}
