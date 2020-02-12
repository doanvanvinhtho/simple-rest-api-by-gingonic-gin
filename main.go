package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/handle"
)

func main() {
	repositoryType := os.Getenv("REPOSITORY_TYPE")
	if len(repositoryType) <= 0 {
		repositoryType = "inmemory"
	}

	handle.Init(repositoryType)
	handle.Setup().Run() // listen and serve on 0.0.0.0:8080
}
