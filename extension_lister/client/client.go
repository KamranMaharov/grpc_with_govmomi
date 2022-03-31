package main

import (
	"io"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "dummymodule/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:3030",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("No way. Couldn't connect to local grpc server. err: %v", err)
	}
	defer conn.Close()

	c := pb.NewExtensionListerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	authResponse, err := c.Authenticate(ctx,
			&pb.AuthRequest{Username: "admin", Password: "admin"})
	if err != nil {
		log.Fatalf("error not expected in auth request: %v", err)
	}
	log.Printf("Auth request completed. response: %v", authResponse)

	stream, err := c.ListExtensions(ctx, &pb.ListExtensionRequest{KeyContains: "sysmgmt"})
	if err != nil {
		log.Fatalf("error not expected in list_extensions request: %v", err)
	}
	id := 0
	for {
		id += 1
		extension, err := stream.Recv()
		if err == io.EOF {
			log.Print("all extensions fetched in client side")
			break
		}
		if err != nil {
			log.Fatal("grpc error while fetching extensions: %v", err)
		}
		log.Printf("Exception #%d: %v", id, extension)
	}
}
