package mongodb

import (
	"context"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// New an event use MongoDB
func New(mc *mongo.Client) repository.Event {
	return eventMongoDB{
		mongoClient: mc,
	}
}

// eventMongoDB is instance of event repository
type eventMongoDB struct {
	mongoClient *mongo.Client
}

// GetOneEvent helps to get a event from MongoDB repository
func (e eventMongoDB) GetOneEvent(id string) (*model.Event, error) {
	var resultEvent model.Event

	collection := e.mongoClient.Database("demo").Collection("event")
	filter := bson.D{{"ID", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&resultEvent)
	if err != nil {
		return nil, err
	}

	return &resultEvent, nil
}
