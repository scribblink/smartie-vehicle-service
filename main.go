package main

import (
	"context"
	"fmt"
	"os"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/scribblink/smartie-vehicle-service/proto/vehicle"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	srv := micro.NewService(
		micro.Name("shippy.service.vehicle"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	vehicleCollection := client.Database("shippy").Collection("vehicle")
	repository := &VehicleRepository{
		vehicleCollection,
	}


	// Register our implementation with
	pb.RegisterVehicleServiceHandler(srv.Server(), &handler{repository})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
