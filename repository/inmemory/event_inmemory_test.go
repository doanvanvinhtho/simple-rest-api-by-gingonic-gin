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
		event, err := e.GetOneEvent(c.inID)

		assert.Nil(t, err)
		assert.Equal(t, c.wantTitle, event.Title)
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
