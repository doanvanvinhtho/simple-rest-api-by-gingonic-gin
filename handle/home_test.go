package handle

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/stretchr/testify/assert"
)

func TestHomeLink(t *testing.T) {
	router := Setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	// Check HTTP response code
	assert.Equal(t, 200, w.Code)

	// Convert the JSON response to object
	var response model.Response
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Welcome home!", response.Message)
}
