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
docker build -t doanvanvinhtho/simple-rest-api-by-gingonic-gin:latest .
docker container run --detach --publish 8080:8080 doanvanvinhtho/simple-rest-api-by-gingonic-gin
```