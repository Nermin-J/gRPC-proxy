package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "grpc-proxy/grpc-app/stubs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type rideServiceServer struct{
	pb.UnimplementedRideServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRideServiceServer(grpcServer, &rideServiceServer{})
	reflection.Register(grpcServer)

	log.Println("gRPC server running on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *rideServiceServer) RequestRide(ctx context.Context, req *pb.RideRequest) (*pb.RideResponse, error) {
	// Simulate driver assignment
	driverID := fmt.Sprintf("driver-%d", rand.Intn(100))

	return &pb.RideResponse{
		RideId:    fmt.Sprintf("ride-%d", rand.Intn(10000)),
		DriverId:  driverID,
		DriverName: "John Doe",
		CarModel:  "Toyota Prius",
		CarPlate:  fmt.Sprintf("ABC-%d", rand.Intn(1000)),
		Status:    "assigned",
	}, nil
}

func (s *rideServiceServer) StreamDriverLocation(req *pb.DriverLocationRequest, stream pb.RideService_StreamDriverLocationServer) error {
	for i := 0; i < 10; i++ {
		loc := &pb.DriverLocation{
			DriverId: req.RideId,
			Latitude:  37.7749 + rand.Float64()*0.01, // rand coordinates
			Longitude: -122.4194 + rand.Float64()*0.01,
			Timestamp: time.Now().Format(time.RFC3339),
		}
		if err := stream.Send(loc); err != nil {
			return err
		}
		time.Sleep(time.Second) // Simulate real-time updates
	}
	return nil
}

func (s *rideServiceServer) UpdateDriverLocation(stream pb.RideService_UpdateDriverLocationServer) error {
	for {
		loc, err := stream.Recv()
		if err != nil {
			log.Printf("Driver disconnected: %v", err)
			break
		}
		log.Printf("Received location update from %s: (%f, %f)", loc.DriverId, loc.Latitude, loc.Longitude)
	}
	return stream.SendAndClose(&pb.LocationUpdateResponse{Status: "success"})
}

func (s *rideServiceServer) CompleteRide(ctx context.Context, req *pb.RideCompletionRequest) (*pb.RideCompletionResponse, error) {
	return &pb.RideCompletionResponse{
		Status:        "completed",
		PaymentStatus: "paid",
	}, nil
}
