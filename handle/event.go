package handle

import (
	"net/http"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository"
	"github.com/gin-gonic/gin"
)

// GetOneEvent helps to get an event from repository
func GetOneEvent(c *gin.Context) {
	var response model.Response

	// Get the ID from the url
	eventID := c.Param("id")

	if event, err := repository.GetOneEvent(eventID); err != nil {
		response = model.Response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		}
	} else {
		response = model.Response{
			Code:    http.StatusOK,
			Message: "",
			Data:    *event,
		}
	}

	c.JSON(http.StatusOK, response)
}
