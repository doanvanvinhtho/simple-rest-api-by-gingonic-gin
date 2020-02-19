package service

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository/inmemory"
)

func TestGetOneEvent(t *testing.T) {
	serviceEvent := New(inmemory.New())

	response := serviceEvent.GetOneEvent("id_gin")
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.Code)

	response = serviceEvent.GetOneEvent("id_not_exist")
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusNotFound, response.Code)
}

func TestGetAllEvent(t *testing.T) {
}
