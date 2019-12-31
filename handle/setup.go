package handle

import (
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/service"
	"github.com/gin-gonic/gin"
)

var serviceEvent service.Event

// Init services
func Init() {
	serviceEvent = service.New(repository.New())
}

// Setup handlers
func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/events/:id", GetOneEvent)

	return r
}
