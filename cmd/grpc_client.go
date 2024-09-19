package main

import (
	"context"
	"log"
	"time"

	pb "path/to/protobufs"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewADRServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ADRRequest{Adr: "Sample ADR Data"}
	res, err := client.AnalyzeADR(ctx, req)
	if err != nil {
		log.Fatalf("Could not analyze ADR: %v", err)
	}
	log.Printf("ADR Analysis Result: %s", res.Result)
}
