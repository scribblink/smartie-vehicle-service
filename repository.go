package main

import (
	"context"
	pb "github.com/scribblink/smartie-vehicle-service/proto/vehicle"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type repository interface {
	FindAvailable(spec *pb.Specification) (*pb.Vehicle, error)
	Create(vehicle *pb.Vehicle) error
}

type VehicleRepository struct {
	collection *mongo.Collection
}

// FindAvailable - checks a specification against a map of vehicles,
// if capacity and max weight are below a vehicles capacity and max weight,
// then return that vehicle.
func (repository *VehicleRepository) FindAvailable(spec *pb.Specification) (*pb.Vehicle, error) {
	filter := bson.D{{
		"capacity",
		bson.D{{
			"$lte",
			spec.Capacity,
		}, {
			"$lte",
			spec.MaxWeight,
		}},
	}}
	var vehicle *pb.Vehicle
	if err := repository.collection.FindOne(context.TODO(), filter).Decode(&vehicle); err != nil {
		return nil, err
	}
	return vehicle, nil
}

// Create a new vehicle
func (repository *VehicleRepository) Create(vehicle *pb.Vehicle) error {
	_, err := repository.collection.InsertOne(context.TODO(), vehicle)
	return err
}
