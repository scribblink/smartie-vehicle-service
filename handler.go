package main

import (
	"context"
	pb "github.com/scribblink/smartie-vehicle-service/proto/vehicle"
)

type handler struct {
	repository
}

// FindAvailable vehicles
func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	// Find the next available vehicle
	vehicle, err := s.repository.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vehicle as part of the response message type
	res.Vehicle = vehicle
	return nil
}

// Create a new vehicle
func (s *handler) Create(ctx context.Context, req *pb.Vehicle, res *pb.Response) error {
	if err := s.repository.Create(req); err != nil {
		return err
	}
	res.Vehicle = req
	return nil
}
