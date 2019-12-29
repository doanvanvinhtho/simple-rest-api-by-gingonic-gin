package handle

import (
	"net/http"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/gin-gonic/gin"
)

// HomeLink shows welcome message
func HomeLink(c *gin.Context) {
	response := model.Response{
		Code:    http.StatusOK,
		Message: "Welcome home!",
		Data:    nil,
	}

	c.JSON(http.StatusOK, response)
}
