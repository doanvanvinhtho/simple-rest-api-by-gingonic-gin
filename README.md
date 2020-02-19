# simple-rest-api-by-gingonic-gin
Learn how to build a simple RESTful API by Go, Just for fun!

## 1. Prerequisites
* Have Go installed

## 2. Test
``` bash
go test ./...
```

## 3. Run the HTTP server
``` bash
go run main.go
```

## 4. Test APIs

### 4.1. Get event by ID
``` bash
curl http://localhost:8080/events/id_go
{"code":200,"data":{"id":"id_go","title":"Go","description":"https://golang.org/"}}

curl http://localhost:8080/events/id_gin
{"code":200,"data":{"id":"id_gin","title":"Gin","description":"https://github.com/gin-gonic/gin"}}

curl http://localhost:8080/events/id_not_exist
{"code":404,"message":"Event id_not_exist not found"}
```

### 4.2. Get all events
``` bash
curl http://localhost:8080/events
{"code":200,"data":[{"id":"id_go","title":"Go","description":"https://golang.org/"},
{"id":"id_gin","title":"Gin","description":"https://github.com/gin-gonic/gin"}]}
```

### 4.3. Add an event
``` bash
curl -H "Content-Type: application/json" -X POST \
     -d '{"id":"a","title":"b","description":"c"}' http://localhost:8080/events
{"code":200}
```

### 4.4. Update an event
``` bash
curl -H "Content-Type: application/json" -X PATCH \
     -d '{"id":"a","title":"b1","description":"c1"}' http://localhost:8080/events
{"code":200}
```

### 4.5. Delete an event
``` bash
curl -X DELETE http://localhost:8080/events/id_go
{"code":200}
```

### 4.6. Stress-Test APIs
``` bash
for i in $(seq 5000); do curl http://localhost:8080/events; done
```

## 5. Docker
``` bash
docker-compose up
```
