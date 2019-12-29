package handle

import "github.com/gin-gonic/gin"

// Setup handlers
func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/", HomeLink)
	r.GET("/events/:id", GetOneEvent)

	return r
}
