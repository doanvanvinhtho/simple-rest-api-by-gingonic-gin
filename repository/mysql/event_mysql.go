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

// GetOneEvent helps to get a event from repository
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

// GetAllEvent helps to get all events from repository
func (e *eventMySQL) GetAllEvent() ([]model.Event, error) {
	return nil, nil
}

// AddEvent helps to add new event into repository
func (e *eventMySQL) AddEvent(ev *model.Event) (string, error) {
	if ev == nil {
		return "", errors.New("[MySQL] The event need to add is nil")
	}
	return "", nil
}

// UpdateEvent helps to update an event in repository
func (e *eventMySQL) UpdateEvent(ev *model.Event) (string, error) {
	if ev == nil {
		return "", errors.New("[MySQL] The event need to update is nil")
	}
	return "", nil
}

// DeleteEvent helps to delete an event in repository
func (e *eventMySQL) DeleteEvent(id string) error {
	if len(id) <= 0 {
		return errors.New("[MySQL] The event id is empty")
	}
	return nil
}
