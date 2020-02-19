# simple-rest-api-by-gingonic-gin
Learn how to build a simple RESTful API by Go, Just for fun!

## Prerequisites
* Have Go installed

## Test
``` bash
go test ./...
```

## Run the HTTP server
``` bash
go run main.go
```

## Test APIs

### Get
``` bash
curl http://localhost:8080/events/id_go
{"code":200,"data":{"id":"id_go","title":"Go","description":"https://golang.org/"}}

curl http://localhost:8080/events/id_gin
{"code":200,"data":{"id":"id_gin","title":"Gin","description":"https://github.com/gin-gonic/gin"}}

curl http://localhost:8080/events/id_not_exist
{"code":404,"message":"Event id_not_exist not found"}
```

### GetAll
``` bash
curl http://localhost:8080/events
{"code":200,"data":[{"id":"id_go","title":"Go","description":"https://golang.org/"},
{"id":"id_gin","title":"Gin","description":"https://github.com/gin-gonic/gin"}]}
```

### Add
``` bash
curl -H "Content-Type: application/json" -X POST \
     -d '{"id":"a","title":"b","description":"c"}' http://localhost:8080/events
{"code":200}
```

### Update
``` bash
curl -H "Content-Type: application/json" -X PATCH \
     -d '{"id":"a","title":"b1","description":"c1"}' http://localhost:8080/events
{"code":200}
```

### Delete
``` bash
curl -X DELETE http://localhost:8080/events/id_go
{"code":200}
```

### Stress-Test
``` bash
for i in $(seq 5000); do curl http://localhost:8080/events; done
```

## Docker
``` bash
docker-compose up
```
