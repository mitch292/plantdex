package main

import (
	"log"
	"net"

	"github.com/mitch292/plantdex/plants"
	"google.golang.org/grpc"
)

func main() {

	plants.InitDB()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	plants.RegisterPlantsServiceServer(grpcServer, &plants.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}
