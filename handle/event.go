package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOneEvent helps to get an event
func GetOneEvent(c *gin.Context) {
	c.JSON(http.StatusOK, serviceEvent.GetOneEvent(c.Param("id")))
}
