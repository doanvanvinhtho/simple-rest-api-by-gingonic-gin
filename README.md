# simple-rest-api-by-gingonic-gin
Learn how to build a simple RESTful API by Go

## Prerequisites
* Have Go installed

## Test
```
go test ./...
```

## Run the HTTP server
```
go run main.go
```

## Test APIs
```
curl http://localhost:8080/events/id_go
{"code":200,"data":{"id":"id_go","title":"Go","description":"https://golang.org/"}}

curl http://localhost:8080/events/id_gin
{"code":200,"data":{"id":"id_gin","title":"Gin","description":"https://github.com/gin-gonic/gin"}}

curl http://localhost:8080/events/12345_54321
{"code":404,"message":"Event 12345_54321 not found"}
```

## Docker
```
docker-compose up
```

### Insert sample data into Redis
```
docker container exec -it simple-rest-api-by-gingonic-gin_redis_1 redis-cli
HMSET event:id_go ID id_go Title "Go" Description "https://golang.org/"
HMSET event:id_gin ID id_gin Title "Gin" Description "https://github.com/gin-gonic/gin"
```

### Insert sample data into MongoDB
```
docker container exec -it simple-rest-api-by-gingonic-gin_mongodb_1 mongo
use demo;
db.event.insert({ID:"id_go", Title:"Go", Description:"https://golang.org/" });
db.event.insert({ID:"id_gin", Title:"Gin", Description:"https://github.com/gin-gonic/gin" });
```