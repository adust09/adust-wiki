package api

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "path/to/protobufs"

	"google.golang.org/grpc"
)

type ADRServer struct {
	pb.UnimplementedADRServiceServer
}

func (s *ADRServer) AnalyzeADR(ctx context.Context, req *pb.ADRRequest) (*pb.ADRResponse, error) {
	result := fmt.Sprintf("Analyzed ADR: %s", req.Adr)
	return &pb.ADRResponse{Result: result}, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterADRServiceServer(s, &ADRServer{})
	log.Println("gRPC server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
