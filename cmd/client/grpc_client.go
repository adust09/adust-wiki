package client

import (
	"context"
	"log"
	"time"

	pb "go-todo/proto/adr"

	"google.golang.org/grpc"
)

func client() {
	// gRPCサーバーに接続
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewADRServiceClient(conn)

	// リクエストを作成
	req := &pb.ADRRequest{Adr: "Sample ADR Data"}

	// サーバーへリクエストを送信
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.AnalyzeADR(ctx, req)
	if err != nil {
		log.Fatalf("Error during request: %v", err)
	}

	// 結果を表示
	log.Printf("Response from server: %s", res.GetResult())
}
