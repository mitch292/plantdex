package plants

import (
	"log"

	"golang.org/x/net/context"
)

// Server is the gRPC server struct for interacting with our plants
type Server struct {
}

// GetPlant is a function used by our gRPC server to return plant by a given name
func (s *Server) GetPlant(ctx context.Context, rp *RequestPlant) (*Plant, error) {
	plant, err := getPlantFromDB(rp.ID)
	if err != nil {
		log.Fatalf("Failure fetching the plant from the DB: %s\n", err)
	}

	return plant, nil
}

// GetAllPlants is a function used by our gRPC server to return plant by a given name
func (s *Server) GetAllPlants(ctx context.Context, e *Empty) (*Plants, error) {
	plants, err := getAllPlantsFromDB()
	if err != nil {
		log.Fatalf("Failure fetching all the plants from the DB: %s\n", err)
	}

	return plants, nil
}

// AddPlant will add a new plant to our database
func (s *Server) AddPlant(ctx context.Context, plant *Plant) (*Feedback, error) {
	_, err := addPlantToDB(plant)
	if err != nil {
		log.Fatalf("Failure adding this plant to the DB: %s\n", err)
	}

	return &Feedback{Success: true, Message: "added"}, nil
}

// UpdatePlant will update an existing plant in the DB
func (s *Server) UpdatePlant(ctx context.Context, plant *Plant) (*Feedback, error) {
	_, err := updatePlantInDB(plant)
	if err != nil {
		log.Fatalf("Failure adding this plant to the DB: %s\n", err)
	}

	return &Feedback{Success: true, Message: "updated"}, nil
}

// DeletePlant is a function used to remove a plant from our database
func (s *Server) DeletePlant(ctx context.Context, rp *RequestPlant) (*Feedback, error) {
	_, err := deletePlantFromDB(rp.ID)
	if err != nil {
		log.Fatalf("Failure deleting the plant from the DB: %s\n", err)
	}

	return &Feedback{Success: true, Message: "removed"}, nil
}
