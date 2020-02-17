package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
)

// GetOneEvent helps to get an event
// It always returns http code 200
func GetOneEvent(c *gin.Context) {
	c.JSON(http.StatusOK, serviceEvent.GetOneEvent(c.Param("id")))
}

// GetAllEvent helps to get all events
// It always returns http code 200
func GetAllEvent(c *gin.Context) {
	c.JSON(http.StatusOK, serviceEvent.GetAllEvent())
}

// AddEvent helps to add new event
// It always returns http code 200
func AddEvent(c *gin.Context) {
	var inp model.Event

	c.BindJSON(&inp)
	c.JSON(http.StatusOK, serviceEvent.AddEvent(&inp))
}

// UpdateEvent helps to update an event
// It always returns http code 200
func UpdateEvent(c *gin.Context) {
	var inp model.Event

	c.BindJSON(&inp)
	c.JSON(http.StatusOK, serviceEvent.UpdateEvent(&inp))
}

// DeleteEvent helps to delete an event
// It always returns http code 200
func DeleteEvent(c *gin.Context) {
	c.JSON(http.StatusOK, serviceEvent.DeleteEvent(c.Param("id")))
}
