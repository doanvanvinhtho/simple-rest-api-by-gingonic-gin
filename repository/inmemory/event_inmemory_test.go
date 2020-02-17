package inmemory

import (
	"testing"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/stretchr/testify/assert"
)

type testcase struct {
	inID            string
	wantTitle       string
	wantDescription string
}

func TestGetOneEventInMemory(t *testing.T) {
	e := New()

	happyCases := []testcase{
		{
			inID:            "id_go",
			wantTitle:       "Go",
			wantDescription: "https://golang.org/",
		},
		{
			inID:            "id_gin",
			wantTitle:       "Gin",
			wantDescription: "https://github.com/gin-gonic/gin",
		},
	}
	for _, c := range happyCases {
		for i := 0; i < 2; i++ {
			event, err := e.GetOneEvent(c.inID)

			assert.Nil(t, err)
			assert.Equal(t, c.wantTitle, event.Title)
			assert.Equal(t, c.wantDescription, event.Description)

			// Try to update the event outside GetOneEvent function
			event.ID = ""
			event.Title = ""
			event.Description = ""
		}
	}

	badCases := []testcase{
		{
			inID:            "12345_54321",
			wantTitle:       "Go_wrong",
			wantDescription: "",
		},
	}
	for _, c := range badCases {
		_, err := e.GetOneEvent(c.inID)

		assert.NotNil(t, err)
	}
}

func TestGetAllEventInMemory(t *testing.T) {
	e := New()

	ev, err := e.GetAllEvent()
	assert.Nil(t, err)
	assert.NotNil(t, ev)
	assert.Equal(t, 2, len(ev))

	happyCases := []testcase{
		{
			inID:            "id_new1",
			wantTitle:       "NewTitle 1",
			wantDescription: "New Description 1",
		},
		{
			inID:            "id_new2",
			wantTitle:       "NewTitle 2",
			wantDescription: "New Description 2",
		},
	}
	for _, c := range happyCases {
		// Add to repository
		event := &model.Event{
			ID:          c.inID,
			Title:       c.wantTitle,
			Description: c.wantDescription,
		}
		id, err := e.AddEvent(event)
		assert.Equal(t, c.inID, id)
		assert.Nil(t, err)
	}

	ev, err = e.GetAllEvent()
	assert.Nil(t, err)
	assert.NotNil(t, ev)
	assert.Equal(t, 2+len(happyCases), len(ev))
}

func TestAddEventInMemory(t *testing.T) {
	e := New()

	happyCases := []testcase{
		{
			inID:            "id_new1",
			wantTitle:       "NewTitle 1",
			wantDescription: "New Description 1",
		},
		{
			inID:            "id_new2",
			wantTitle:       "NewTitle 2",
			wantDescription: "New Description 2",
		},
	}
	for _, c := range happyCases {
		// Get the event, the result is empty
		event, err := e.GetOneEvent(c.inID)
		assert.Nil(t, event)
		assert.NotNil(t, err)

		// Add to repository
		event = &model.Event{
			ID:          c.inID,
			Title:       c.wantTitle,
			Description: c.wantDescription,
		}
		id, err := e.AddEvent(event)
		assert.Equal(t, c.inID, id)
		assert.Nil(t, err)

		// Get the event again to compare
		event, err = e.GetOneEvent(c.inID)
		assert.NotNil(t, event)
		assert.Nil(t, err)
		assert.Equal(t, c.wantTitle, event.Title)
		assert.Equal(t, c.wantDescription, event.Description)
	}
}

func TestUpdateEventInMemory(t *testing.T) {
	e := New()

	happyCases := []testcase{
		{
			inID:            "id_go",
			wantTitle:       "GoNew",
			wantDescription: "GoNew Description",
		},
		{
			inID:            "id_gin",
			wantTitle:       "GinNew",
			wantDescription: "GinNew Description",
		},
	}
	for _, c := range happyCases {
		// Get the event, the result isn't empty
		event, err := e.GetOneEvent(c.inID)
		assert.NotNil(t, event)
		assert.Nil(t, err)

		// Update event's title & description to new value
		event.Title = c.wantTitle
		event.Description = c.wantDescription
		id, err := e.UpdateEvent(event)
		assert.Equal(t, c.inID, id)

		// Get the event again to compare new title
		event, err = e.GetOneEvent(c.inID)
		assert.NotNil(t, event)
		assert.Nil(t, err)
		assert.Equal(t, c.wantTitle, event.Title)
		assert.Equal(t, c.wantDescription, event.Description)
	}

	badCases := []testcase{
		{
			inID:            "12345_54321",
			wantTitle:       "Go_wrong",
			wantDescription: "",
		},
		{
			inID:            "",
			wantTitle:       "",
			wantDescription: "",
		},
	}
	for _, c := range badCases {
		event := &model.Event{
			ID:          c.inID,
			Title:       "",
			Description: "",
		}
		_, err := e.UpdateEvent(event)
		assert.NotNil(t, err)
	}
}

func TestDeleteEventInMemory(t *testing.T) {
	e := New()

	happyCases := []testcase{
		{
			inID:            "id_go",
			wantTitle:       "Go",
			wantDescription: "",
		},
		{
			inID:            "id_gin",
			wantTitle:       "Gin",
			wantDescription: "",
		},
	}
	for _, c := range happyCases {
		// Delete the event
		err := e.DeleteEvent(c.inID)
		assert.Nil(t, err)

		// Try to get the event after delete it
		event, err := e.GetOneEvent(c.inID)
		assert.Nil(t, event)
		assert.NotNil(t, err)
	}

	badCases := []testcase{
		{
			inID:            "12345_54321",
			wantTitle:       "Go_wrong",
			wantDescription: "",
		},
		{
			inID:            "",
			wantTitle:       "",
			wantDescription: "",
		},
	}
	for _, c := range badCases {
		err := e.DeleteEvent(c.inID)
		assert.NotNil(t, err)
	}
}
