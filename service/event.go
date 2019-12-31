package service

import "github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"

// Event service interface
type Event interface {
	GetOneEvent(id string) model.Response
	GetAllEvent() []model.Response
}
