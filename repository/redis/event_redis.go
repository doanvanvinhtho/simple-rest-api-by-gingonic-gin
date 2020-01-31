package redis

import (
	"errors"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository"
	"github.com/gomodule/redigo/redis"
)

// New an event use redis
func New(rp *redis.Pool) repository.Event {
	return eventRedis{
		pool: rp,
	}
}

// eventRedis is instance of event repository
type eventRedis struct {
	pool *redis.Pool
}

// GetOneEvent helps to get a event from redis repository
func (e eventRedis) GetOneEvent(id string) (*model.Event, error) {
	conn := e.pool.Get()
	defer conn.Close()

	// HGETALL: returns all fields and values of the hash stored at key.
	// In the returned value, every field name is followed by its value,
	// so the length of the reply is twice the size of the hash.

	values, err := redis.Values(conn.Do("HGETALL", "event:"+id))
	if err != nil {
		return nil, err
	} else if len(values) <= 0 {
		return nil, errors.New("Event " + id + " not found")
	}

	var resultEvent model.Event
	err = redis.ScanStruct(values, &resultEvent)
	if err != nil {
		return nil, err
	}

	return &resultEvent, nil
}
