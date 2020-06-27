package plants

import (
	"log"

	"golang.org/x/net/context"
)

// Server is the gRPC server struct for interacting with our plants
type Server struct {
}

// GetPlant is a function used by our gRPC server to return plant by a given name
func (s *Server) GetPlant(ctx context.Context, requestPlant *RequestPlant) (*Plant, error) {
	log.Printf("Received a request for a plant: %d", requestPlant.ID)

	plant, err := getPlantFromDB(requestPlant.ID)
	if err != nil {
		log.Fatalf("Failure fetching the plant from the DB: %s\n", err)
		return plant, err
	}

	return plant, nil
}

// AddPlant will add a new plant to our database
func (s *Server) AddPlant(ctx context.Context, plant *Plant) (*Feedback, error) {
	log.Printf("Recieved a request to add a plant: %s", plant)

	id, err := addPlantToDB(plant)
	if err != nil {
		log.Fatalf("Failure adding this plant to the DB: %s\n", err)
	}

	log.Printf("Success %d\n", id)

	return &Feedback{Success: true, Message: "ok"}, nil
}
