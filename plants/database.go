package plants

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Set client options
// TODO: env variable
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")

// Connect to MongoDB
var client, err = mongo.Connect(context.TODO(), clientOptions)

// DB is the connection to our mongo db collection for the plants
// TODO: Env variable
var db = client.Database("plantdex").Collection("plants")

func insertPlant(p Plant) (bool, error) {
	_, err := db.InsertOne(context.TODO(), bson.D{
		{Key: "name", Value: p.name},
		{Key: "size", Value: p.size},
		{Key: "waterSchedule", Value: p.waterSchedule},
		{Key: "sunLevel", Value: p.sunLevel},
		{Key: "notes", Value: p.notes},
		{Key: "isPetSafe", Value: p.isPetSafe},
		{Key: "food", Value: p.food},
		{Key: "shouldMist", Value: p.shouldMist},
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func getPlantByName(name string) (*Plant, error) {
	var foundPlant Plant
	filter := bson.D{{Key: "name", Value: name}}

	err := db.FindOne(context.TODO(), filter).Decode(&foundPlant)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Found a single document: %+v\n", foundPlant)
	return &foundPlant, nil
}

func getAllPlants() ([]*Plant, error) {
	findOptions := options.Find()

	var results []*Plant

	// Grab a cursor
	cur, err := db.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		return results, err
	}

	// Loop through our collection
	for cur.Next(context.TODO()) {
		// create a Plant to be decoded into
		var plant Plant
		err := cur.Decode(&plant)

		if err != nil {
			// return what we have so far and the error
			return results, err
		}

		results = append(results, &plant)

	}

	if err := cur.Err(); err != nil {
		return results, err
	}

	// close the cursor connection
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results, nil

}
