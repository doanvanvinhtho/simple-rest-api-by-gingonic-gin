package redis

import (
	"errors"

	"github.com/gomodule/redigo/redis"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository"
)

// New an event use redis
func New(rp *redis.Pool) repository.Event {
	return &eventRedis{
		pool: rp,
	}
}

// eventRedis is instance of event repository
type eventRedis struct {
	pool *redis.Pool
}

// GetOneEvent helps to get an event from repository
func (e *eventRedis) GetOneEvent(id string) (*model.Event, error) {
	conn := e.pool.Get()
	defer conn.Close()

	// HGETALL: returns all fields and values of the hash stored at key.
	// In the returned value, every field name is followed by its value,
	// so the length of the reply is twice the size of the hash.
	values, err := redis.Values(conn.Do("HGETALL", "event:"+id))
	if err != nil {
		return nil, err
	} else if len(values) <= 0 {
		return nil, errors.New("[Redis] Event " + id + " not found")
	}

	var resultEvent model.Event
	err = redis.ScanStruct(values, &resultEvent)
	if err != nil {
		return nil, err
	}

	return &resultEvent, nil
}

// GetAllEvent helps to get all events from repository
func (e *eventRedis) GetAllEvent() ([]model.Event, error) {
	var resultEvent []model.Event

	conn := e.pool.Get()
	defer conn.Close()

	scanResult, err := redis.MultiBulk(conn.Do("SCAN", 0, "MATCH", "event:*", "COUNT", 20))
	if err != nil {
		return nil, err
	}

	// iterator, err := redis.Int(scanResult[0], nil)
	// if err != nil {
	// 	return nil, err
	// }

	keys, err := redis.Strings(scanResult[1], nil)
	if err != nil {
		return nil, err
	}

	for _, k := range keys {
		values, err := redis.Values(conn.Do("HGETALL", k))
		if err != nil {
			return nil, err
		} else if len(values) <= 0 {
			return nil, errors.New("[Redis] Event " + k + " not found")
		}

		var elmEvent model.Event
		err = redis.ScanStruct(values, &elmEvent)
		if err != nil {
			return nil, err
		}

		resultEvent = append(resultEvent, elmEvent)
	}

	return resultEvent, nil
}

// AddEvent helps to add new event into repository
func (e *eventRedis) AddEvent(ev *model.Event) (string, error) {
	if ev == nil {
		return "", errors.New("[Redis] The event need to add is nil")
	}

	return e.UpdateEvent(ev)
}

// UpdateEvent helps to update an event in repository
func (e *eventRedis) UpdateEvent(ev *model.Event) (string, error) {
	if ev == nil {
		return "", errors.New("[Redis] The event need to update is nil")
	}

	conn := e.pool.Get()
	defer conn.Close()

	// HSET: Sets field in the hash stored at key to value.
	// If key does not exist, a new key holding a hash is created.
	// If field already exists in the hash, it is overwritten
	_, err := conn.Do("HSET", "event:"+ev.ID,
		"ID", ev.ID,
		"Title", ev.Title,
		"Description", ev.Description,
	)
	if err != nil {
		return "", err
	}

	return ev.ID, nil
}

// DeleteEvent helps to delete an event in repository
func (e *eventRedis) DeleteEvent(id string) error {
	if len(id) <= 0 {
		return errors.New("[Redis] The event id is empty")
	}

	conn := e.pool.Get()
	defer conn.Close()

	// DEL: Removes the specified keys. A key is ignored if it does not exist.
	_, err := conn.Do("DEL", "event:"+id)
	if err != nil {
		return err
	}

	return nil
}
