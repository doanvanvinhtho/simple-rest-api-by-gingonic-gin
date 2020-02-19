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
	selCmd, err := e.db.Query("SELECT * FROM event WHERE id=? LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	defer selCmd.Close()

	for selCmd.Next() {
		var id, title, description string
		err = selCmd.Scan(&id, &title, &description)
		if err != nil {
			return nil, err
		}

		resultEvent := model.Event{
			ID:          id,
			Title:       title,
			Description: description,
		}
		return &resultEvent, nil
	}

	return nil, errors.New("[MySQL] Event " + id + " not found")
}

// GetAllEvent helps to get all events from repository
func (e *eventMySQL) GetAllEvent() ([]model.Event, error) {
	var resultEvent []model.Event

	selCmd, err := e.db.Query("SELECT * FROM event LIMIT 20")
	if err != nil {
		return nil, err
	}
	defer selCmd.Close()

	for selCmd.Next() {
		var id, title, description string
		err = selCmd.Scan(&id, &title, &description)
		if err != nil {
			return nil, err
		}

		elem := model.Event{
			ID:          id,
			Title:       title,
			Description: description,
		}
		resultEvent = append(resultEvent, elem)
	}
	if err := selCmd.Err(); err != nil {
		return nil, err
	}

	return resultEvent, nil
}

// AddEvent helps to add new event into repository
func (e *eventMySQL) AddEvent(ev *model.Event) (string, error) {
	if ev == nil {
		return "", errors.New("[MySQL] The event need to add is nil")
	}

	insertCmd, err := e.db.Prepare("INSERT INTO event (id, title, description) VALUES(?, ?, ?)")
	if err != nil {
		return "", err
	}
	defer insertCmd.Close()

	_, err = insertCmd.Exec(ev.ID, ev.Title, ev.Description)
	if err != nil {
		return "", err
	}

	return ev.ID, nil
}

// UpdateEvent helps to update an event in repository
func (e *eventMySQL) UpdateEvent(ev *model.Event) (string, error) {
	if ev == nil {
		return "", errors.New("[MySQL] The event need to update is nil")
	}

	updateCmd, err := e.db.Prepare("UPDATE event SET title=?, description=? WHERE id=?")
	if err != nil {
		return "", err
	}
	defer updateCmd.Close()

	_, err = updateCmd.Exec(ev.Title, ev.Description, ev.ID)
	if err != nil {
		return "", err
	}

	return ev.ID, nil
}

// DeleteEvent helps to delete an event in repository
func (e *eventMySQL) DeleteEvent(id string) error {
	if len(id) <= 0 {
		return errors.New("[MySQL] The event id is empty")
	}

	deleteCmd, err := e.db.Prepare("DELETE FROM event WHERE id=?")
	if err != nil {
		return err
	}
	defer deleteCmd.Close()

	_, err = deleteCmd.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
