package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository"
)

// New an event use MongoDB
func New(mc *mongo.Client) repository.Event {
	return &eventMongoDB{
		mongoClient: mc,
	}
}

// eventMongoDB is instance of event repository
type eventMongoDB struct {
	mongoClient *mongo.Client
}

// GetOneEvent helps to get a event from repository
func (e *eventMongoDB) GetOneEvent(id string) (*model.Event, error) {
	var resultEvent model.Event

	collection := e.mongoClient.Database("demo").Collection("event")
	filter := bson.D{{"ID", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&resultEvent)
	if err != nil {
		return nil, err
	}

	return &resultEvent, nil
}

// GetAllEvent helps to get all events from repository
func (e *eventMongoDB) GetAllEvent() ([]*model.Event, error) {
	return nil, nil
}

// AddEvent helps to add new event into repository
func (e *eventMongoDB) AddEvent(ev *model.Event) (string, error) {
	if ev == nil {
		return "", errors.New("[MongoDB] The event need to add is nil")
	}
	return "", nil
}

// UpdateEvent helps to update an event in repository
func (e *eventMongoDB) UpdateEvent(ev *model.Event) (string, error) {
	if ev == nil {
		return "", errors.New("[MongoDB] The event need to update is nil")
	}
	return "", nil
}

// DeleteEvent helps to delete an event in repository
func (e *eventMongoDB) DeleteEvent(id string) error {
	if len(id) <= 0 {
		return errors.New("[MongoDB] The event id is empty")
	}
	return nil
}
