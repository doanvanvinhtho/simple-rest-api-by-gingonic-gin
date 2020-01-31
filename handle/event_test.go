package handle

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/stretchr/testify/assert"
)

type testcase struct {
	inID     string
	wantCode int
}

func TestGetOneEvent(t *testing.T) {
	Init("inmemory")
	router := Setup()

	cases := []testcase{
		{
			inID:     "id_go",
			wantCode: 200,
		},
		{
			inID:     "12345_54321",
			wantCode: 404,
		},
	}

	for _, c := range cases {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/events/"+c.inID, nil)
		router.ServeHTTP(w, req)

		// The http response code always 200
		assert.Equal(t, 200, w.Code)

		// Convert the JSON response to object
		var response model.Response
		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)

		assert.Equal(t, c.wantCode, response.Code)
	}
}
