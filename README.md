# simple-rest-api-by-gingonic-gin
Learn how to build a simple RESTful API by Go

## Prerequisites
* Have Go installed. If not, check it out [here](https://golang.org/doc/install)
* Also after installing, make sure you are working inside your GOPATH

## Install
```
go get -u github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin
go get -u github.com/gin-gonic/gin
```

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
curl http://localhost:8080/events/6yh4tgrt4y58jdfuer

{"code":200,"data":{"id":"6yh4tgrt4y58jdfuer","title":"Go","description":"https://golang.org/"}}
```