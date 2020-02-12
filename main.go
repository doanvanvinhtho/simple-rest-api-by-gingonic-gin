package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/handle"
)

func main() {
	handle.Init("inmemory")
	handle.Setup().Run() // listen and serve on 0.0.0.0:8080
}
