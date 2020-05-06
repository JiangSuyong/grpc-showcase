package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/longkai/grpc-showcase/genproto/apis/v1"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:50051", "server addr")
	name = flag.String("name", "world", "name argument")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
