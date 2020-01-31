package service

import (
	"net/http"
	"testing"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestGetOneEvent(t *testing.T) {
	serviceEvent := New(inmemory.New())

	response := serviceEvent.GetOneEvent("id_gin")
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.Code)

	response = serviceEvent.GetOneEvent("12345_54321")
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusNotFound, response.Code)
}

func TestGetAllEvent(t *testing.T) {
}
