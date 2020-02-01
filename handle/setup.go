package handle

import (
	"context"
	"log"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository/inmemory"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository/mongodb"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository/mysql"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository/redis"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/service"
	"github.com/gin-gonic/gin"
	redisDriver "github.com/gomodule/redigo/redis"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

var serviceEvent service.Event

func initInMemoryRepo() {
	serviceEvent = service.New(inmemory.New())
}

func initMongoDBRepo() {
	// Set client options
	clientOptions := mongoOptions.Client().ApplyURI("mongodb://localhost:27017/demo")
	clientOptions.SetMaxPoolSize(50)

	// Connect to MongoDB
	client, err := mongoDriver.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	serviceEvent = service.New(mongodb.New(client))
}

func initMySQLRepo() {
	serviceEvent = service.New(mysql.New())
}

func initRedisRepo() {
	redisPool := &redisDriver.Pool{
		// Maximum number of connections allocated by the pool at a given time.
		// When zero, there is no limit on the number of connections in the pool.
		MaxActive: 20,
		// Maximum number of idle connections in the pool.
		MaxIdle: 10,
		// Close connections after remaining idle for this duration. If the value
		// is zero, then idle connections are not closed. Applications should set
		// the timeout to a value less than the server's timeout.
		IdleTimeout: 0,
		// If Wait is true and the pool is at the MaxActive limit, then Get() waits
		// for a connection to be returned to the pool before returning.
		Wait: true,
		Dial: func() (redisDriver.Conn, error) {
			return redisDriver.Dial("tcp", "localhost:6379")
		},
	}
	serviceEvent = service.New(redis.New(redisPool))
}

// Init services
func Init(repo string) {
	switch repo {
	case "inmemory":
		initInMemoryRepo()
	case "mongodb":
		initMongoDBRepo()
	case "mysql":
		initMySQLRepo()
	case "redis":
		initRedisRepo()
	}
}

// Setup handlers
func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/events/:id", GetOneEvent)

	return r
}
