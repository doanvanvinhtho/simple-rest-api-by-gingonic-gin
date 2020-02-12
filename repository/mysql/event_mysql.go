package mysql

import (
	"errors"

	"database/sql"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository"
)

// New an event use MySQL
func New(db *sql.DB) repository.Event {
	return &eventMySQL{
		db: db,
	}
}

// eventMySQL is instance of event repository
type eventMySQL struct {
	db *sql.DB
}

// GetOneEvent helps to get a event from MySQL repository
func (e *eventMySQL) GetOneEvent(id string) (*model.Event, error) {
	selCmd, err := e.db.Query("SELECT * FROM event WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	resultEvent := model.Event{}
	for selCmd.Next() {
		var id, title, description string
		err = selCmd.Scan(&id, &title, &description)
		if err != nil {
			return nil, err
		}
		resultEvent.ID = id
		resultEvent.Title = title
		resultEvent.Description = description

		return &resultEvent, nil
	}

	return nil, errors.New("[MySQL] Event " + id + " not found")
}
