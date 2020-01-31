package handle

import (
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository/inmemory"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository/mongodb"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository/mysql"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository/redis"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/service"
	"github.com/gin-gonic/gin"
)

var serviceEvent service.Event

// Init services
func Init(repo string) {
	switch repo {
	case "inmemory":
		serviceEvent = service.New(inmemory.New())
	case "mongodb":
		serviceEvent = service.New(mongodb.New())
	case "mysql":
		serviceEvent = service.New(mysql.New())
	case "redis":
		serviceEvent = service.New(redis.New())
	}
}

// Setup handlers
func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/events/:id", GetOneEvent)

	return r
}
