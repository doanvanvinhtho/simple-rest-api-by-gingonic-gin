package repository

import "testing"

type testcase struct {
	inID      string
	wantTitle string
}

func TestGetOneEvent(t *testing.T) {
	happyCases := []testcase{
		{
			inID:      "6yh4tgrt4y58jdfuer",
			wantTitle: "Go",
		},
		{
			inID:      "uyhetgdtrg36223jye",
			wantTitle: "Gin",
		},
	}
	for _, c := range happyCases {
		if event, err := GetOneEvent(c.inID); err != nil {
			t.Errorf("GetOneEvent has error NOT FOUND with ID %v", c.inID)
		} else if event.Title != c.wantTitle {
			t.Errorf("GetOneEvent has error WRONG TITLE with ID %v (want: %v, get: %v)", c.inID, c.wantTitle, event.Title)
		}
	}

	badCases := []testcase{
		{
			inID:      "1234554321",
			wantTitle: "Go_wrong",
		},
	}
	for _, c := range badCases {
		if _, err := GetOneEvent(c.inID); err == nil {
			t.Errorf("GetOneEvent has error FOUND with ID %v", c.inID)
		}
	}
}
