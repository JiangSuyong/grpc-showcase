package main

import (
	"context"
	"flag"
	"os"
	"time"

	pb "github.com/longkai/grpc-showcase/genproto/apis/v1"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
)

var (
	addr = flag.String("addr", "localhost:50051", "server addr, e.g., ip:port")
	name = flag.String("name", "world", "name argument")
)

func main() {
	klog.InitFlags(nil)
	defer klog.Flush()

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		klog.ErrorS(err, "Could not connect")
		os.Exit(1)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		klog.ErrorS(err, "Could not greet")
		os.Exit(1)
	}
	klog.InfoS("Greet", "msg", r.Message)
}
