package main

import (
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/handle"
)

func main() {
	handle.Setup().Run() // listen and serve on 0.0.0.0:8080
}
