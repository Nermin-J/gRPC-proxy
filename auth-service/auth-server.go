package main

import (
	"context"
	"fmt"
	"log"
	"net"

	ext_authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type authServer struct {
	ext_authz.UnimplementedAuthorizationServer
}

func (s *authServer) Check(ctx context.Context, req *ext_authz.CheckRequest) (*ext_authz.CheckResponse, error) {
	// Extract headers from the request
	headers := req.Attributes.Request.Http.Headers
	fmt.Println(headers)

	// Example: Check for an "Authorization" header
	authHeader, ok := headers["authorization"]

	if !ok || authHeader == "" {
		log.Println("Authorization header missing")
		return &ext_authz.CheckResponse{
			Status: &status.Status{
				Code:    int32(codes.Unauthenticated),
				Message: "Authorization header missing",
			},
		}, nil
	}

	// Example: Simple authorization logic (replace with your actual logic)
	if authHeader == "Bearer valid_token" { // example token check
		log.Println("Authorization successful")
		return &ext_authz.CheckResponse{
			Status: &status.Status{
				Code: int32(codes.OK),
			},
		}, nil
	}

	log.Println("Authorization failed")
	return &ext_authz.CheckResponse{
		Status: &status.Status{
			Code:    int32(codes.Unauthenticated),
			Message: "Invalid authorization token",
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:40041")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ext_authz.RegisterAuthorizationServer(s, &authServer{})

	fmt.Println("Auth service listening on 0.0.0.0:40041")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
