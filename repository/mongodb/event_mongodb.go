package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/model"
	"github.com/doanvanvinhtho/simple-rest-api-by-gingonic-gin/repository"
)

// New an event use MongoDB
func New(mc *mongo.Client, db *mongo.Database) repository.Event {
	return &eventMongoDB{
		mongoClient: mc,
		mongoDB:     db,
	}
}

// eventMongoDB is instance of event repository
type eventMongoDB struct {
	mongoClient *mongo.Client
	mongoDB     *mongo.Database
}

// GetOneEvent helps to get a event from repository
func (e *eventMongoDB) GetOneEvent(id string) (*model.Event, error) {
	var resultEvent model.Event

	collection := e.mongoDB.Collection("event")
	filter := bson.D{{"id", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&resultEvent)
	if err != nil {
		return nil, err
	}

	return &resultEvent, nil
}

// GetAllEvent helps to get all events from repository
func (e *eventMongoDB) GetAllEvent() ([]model.Event, error) {
	var resultEvent []model.Event

	collection := e.mongoDB.Collection("event")

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(20)

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}
	// Close the cursor once finished
	defer cur.Close(context.TODO())

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// Create a value into which the single document can be decoded
		var elem model.Event
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		resultEvent = append(resultEvent, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return resultEvent, nil
}

// AddEvent helps to add new event into repository
func (e *eventMongoDB) AddEvent(ev *model.Event) (string, error) {
	if ev == nil {
		return "", errors.New("[MongoDB] The event need to add is nil")
	}

	collection := e.mongoDB.Collection("event")
	_, err := collection.InsertOne(context.TODO(), ev)
	if err != nil {
		return "", err
	}

	return ev.ID, nil
}

// UpdateEvent helps to update an event in repository
func (e *eventMongoDB) UpdateEvent(ev *model.Event) (string, error) {
	if ev == nil {
		return "", errors.New("[MongoDB] The event need to update is nil")
	}

	collection := e.mongoDB.Collection("event")
	filter := bson.D{{"id", ev.ID}}
	update := bson.M{
		"$set": bson.M{
			"title":       ev.Title,
			"description": ev.Description,
		}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "", err
	}
	if updateResult.MatchedCount != 1 {
		return "", errors.New("[MongoDB] The event " + ev.ID + " was not found")
	}

	return ev.ID, nil
}

// DeleteEvent helps to delete an event in repository
func (e *eventMongoDB) DeleteEvent(id string) error {
	if len(id) <= 0 {
		return errors.New("[MongoDB] The event id is empty")
	}

	collection := e.mongoDB.Collection("event")
	filter := bson.D{{"id", id}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount != 1 {
		return errors.New("[MongoDB] The event " + id + " was not found")
	}

	return nil
}
