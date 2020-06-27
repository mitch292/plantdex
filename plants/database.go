package plants

import (
	"context"
	"log"

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
		{Key: "name", Value: p.Name},
		{Key: "size", Value: p.Size},
		{Key: "waterSchedule", Value: p.WaterSchedule},
		{Key: "sunLevel", Value: p.SunLevel},
		{Key: "notes", Value: p.Notes},
		{Key: "isPetSafe", Value: p.IsPetSafe},
		{Key: "food", Value: p.Food},
		{Key: "shouldMist", Value: p.ShouldMist},
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func getPlantByName(name string) (*Plant, error) {
	var foundPlant Plant
	filter := bson.M{"name": name}

	err := db.FindOne(context.TODO(), filter).Decode(&foundPlant)

	if err != nil {
		return nil, err
	}

	return &foundPlant, nil
}

func getAllPlants() ([]Plant, error) {

	cursor, err := db.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var plants []Plant

	if err = cursor.All(context.TODO(), &plants); err != nil {
		return plants, err
	}

	return plants, nil
}
