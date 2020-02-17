package inmemory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	inID      string
	wantTitle string
}

func TestGetOneEventInMemory(t *testing.T) {
	e := New()

	happyCases := []testcase{
		{
			inID:      "id_go",
			wantTitle: "Go",
		},
		{
			inID:      "id_gin",
			wantTitle: "Gin",
		},
	}
	for _, c := range happyCases {
		for i := 0; i < 2; i++ {
			event, err := e.GetOneEvent(c.inID)

			assert.Nil(t, err)
			assert.Equal(t, c.wantTitle, event.Title)

			// Try to update the event outside GetOneEvent function
			event.ID = ""
			event.Title = ""
			event.Description = ""
		}
	}

	badCases := []testcase{
		{
			inID:      "12345_54321",
			wantTitle: "Go_wrong",
		},
	}
	for _, c := range badCases {
		_, err := e.GetOneEvent(c.inID)

		assert.NotNil(t, err)
	}
}

func TestDeleteEventInMemory(t *testing.T) {
	e := New()

	happyCases := []testcase{
		{
			inID:      "id_go",
			wantTitle: "Go",
		},
		{
			inID:      "id_gin",
			wantTitle: "Gin",
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
			inID:      "12345_54321",
			wantTitle: "Go_wrong",
		},
		{
			inID:      "",
			wantTitle: "",
		},
	}
	for _, c := range badCases {
		err := e.DeleteEvent(c.inID)
		assert.NotNil(t, err)
	}
}
