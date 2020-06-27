package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/mitch292/plantdex/plants"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client could not connect: %s\n", err)
	}
	defer conn.Close()

	p := plants.NewPlantsServiceClient(conn)

	// Test requesting a plant
	plant := plants.RequestPlant{ID: 1}

	response, err := p.GetPlant(context.Background(), &plant)
	if err != nil {
		log.Fatalf("Error calling GetPlant: %s\n", err)
	}

	log.Printf("Response from the server: %s\n", response)

	// Test adding a plant
	newPlant := plants.Plant{
		Name:          "System Added",
		Size:          2,
		WaterSchedule: 400532,
		SunLevel:      3,
		Notes:         "This was added via our gRPC server, pretty cool.",
		IsPetSafe:     false,
		Food:          9432990,
		ShouldMist:    true,
	}

	addResp, err := p.AddPlant(context.Background(), &newPlant)
	if err != nil {
		log.Fatalf("Error Adding a plant: %s\n", err)
	}

	log.Printf("We added a plant!!: %v\n", addResp)

}