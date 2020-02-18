package handle

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
)

type testcase struct {
	inID            string
	inTitle         string
	inDescription   string
	wantCode        int
	wantTitle       string
	wantDescription string
}

func TestGetOneEvent(t *testing.T) {
	Init("inmemory")
	router := Setup()

	cases := []testcase{
		{
			inID:            "id_go",
			inTitle:         "Go",
			inDescription:   "https://golang.org/",
			wantCode:        200,
			wantTitle:       "Go",
			wantDescription: "https://golang.org/",
		},
		{
			inID:            "12345_54321",
			inTitle:         "",
			inDescription:   "",
			wantCode:        404,
			wantTitle:       "",
			wantDescription: "",
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

		if response.Code == 200 {
			resData := response.Data.(map[string]interface{})
			assert.Equal(t, c.wantTitle, resData["title"])
			assert.Equal(t, c.wantDescription, resData["description"])
		}
	}
}

func TestGetAllEvent(t *testing.T) {
	Init("inmemory")
	router := Setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/events", nil)
	router.ServeHTTP(w, req)

	// The http response code always 200
	assert.Equal(t, 200, w.Code)

	// Convert the JSON response to object
	var response model.Response
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)

	assert.Equal(t, 200, response.Code)
}

func TestAddEvent(t *testing.T) {
	Init("inmemory")
	router := Setup()

	cases := []testcase{
		{
			inID:            "id_new_1",
			inTitle:         "title_new_1",
			inDescription:   "description_new_1",
			wantCode:        200,
			wantTitle:       "title_new_1",
			wantDescription: "description_new_1",
		},
	}

	for _, c := range cases {
		w := httptest.NewRecorder()

		reqEventJSON, _ := json.Marshal(model.Event{
			ID:          c.inID,
			Title:       c.inTitle,
			Description: c.inDescription,
		})
		req, _ := http.NewRequest("POST", "/events", bytes.NewBuffer([]byte(reqEventJSON)))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		// The http response code always 200
		assert.Equal(t, 200, w.Code)

		// Convert the JSON response to object
		var response model.Response
		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)
		assert.Equal(t, c.wantCode, response.Code)

		// Get the event back to check
		w = httptest.NewRecorder()

		req, _ = http.NewRequest("GET", "/events/"+c.inID, nil)
		router.ServeHTTP(w, req)

		// The http response code always 200
		assert.Equal(t, 200, w.Code)

		// Convert the JSON response to object
		err = json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)
		assert.Equal(t, c.wantCode, response.Code)

		if response.Code == 200 {
			resData := response.Data.(map[string]interface{})
			assert.Equal(t, c.wantTitle, resData["title"])
			assert.Equal(t, c.wantDescription, resData["description"])
		}
	}
}

func TestUpdateEvent(t *testing.T) {
	Init("inmemory")
	router := Setup()

	cases := []testcase{
		{
			inID:            "id_go",
			inTitle:         "title_new_1",
			inDescription:   "description_new_1",
			wantCode:        200,
			wantTitle:       "title_new_1",
			wantDescription: "description_new_1",
		},
	}

	for _, c := range cases {
		w := httptest.NewRecorder()

		reqEventJSON, _ := json.Marshal(model.Event{
			ID:          c.inID,
			Title:       c.inTitle,
			Description: c.inDescription,
		})
		req, _ := http.NewRequest("PATCH", "/events", bytes.NewBuffer([]byte(reqEventJSON)))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		// The http response code always 200
		assert.Equal(t, 200, w.Code)

		// Convert the JSON response to object
		var response model.Response
		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)
		assert.Equal(t, c.wantCode, response.Code)

		// Get back the event
		w = httptest.NewRecorder()

		req, _ = http.NewRequest("GET", "/events/"+c.inID, nil)
		router.ServeHTTP(w, req)

		// The http response code always 200
		assert.Equal(t, 200, w.Code)

		// Convert the JSON response to object
		err = json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)
		assert.Equal(t, c.wantCode, response.Code)

		if response.Code == 200 {
			resData := response.Data.(map[string]interface{})
			assert.Equal(t, c.wantTitle, resData["title"])
			assert.Equal(t, c.wantDescription, resData["description"])
		}
	}
}

func TestDeleteEvent(t *testing.T) {
	Init("inmemory")
	router := Setup()

	cases := []testcase{
		{
			inID:            "id_go",
			inTitle:         "",
			inDescription:   "",
			wantCode:        404,
			wantTitle:       "",
			wantDescription: "",
		},
	}

	for _, c := range cases {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/events/"+c.inID, nil)
		router.ServeHTTP(w, req)

		// The http response code always 200
		assert.Equal(t, 200, w.Code)

		// Get the event to make sure it doesn't exist
		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/events/"+c.inID, nil)
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
