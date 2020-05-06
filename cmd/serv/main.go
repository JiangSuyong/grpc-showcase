package main

import (
	"context"
	"flag"
	grpc_showcase "github.com/longkai/grpc-showcase"
	pb "github.com/longkai/grpc-showcase/genproto/apis/v1"
	pbLibrary "github.com/longkai/grpc-showcase/genproto/apis/library/v1"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("%+v", req)
	// authentication (token verification)
	m, err := handler(ctx, req)
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}
	return m, err
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v, age: %d", in.GetName(), in.GetAge())
	md, ok := metadata.FromIncomingContext(ctx)
	log.Println(md, ok)
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// CreateXml _
func (s *server) CreateXml(ctx context.Context, in *pb.XmlRequest) (*httpbody.HttpBody, error) {
	if true {
		return nil, status.Error(codes.NotFound, "not found")
	}
	log.Printf("x-req-id %s, ct %s, body: %s", in.GetRequestId(), in.HttpBody.ContentType, in.HttpBody.Data)
	res := httpbody.HttpBody{
		ContentType: "application/xml; charset=utf-8",
		Data: []byte(`<xml>
    <AppId><![CDATA[wxf1c1a8cfdbda189d]]></AppId>
    <Encrypt><![CDATA[MB837iWog1s3Xsvk3mA6qJIYiG9emT5g0yPUml3oKnnUA5a+9WhIFM/QVENoBJsR2fTi6bzaM51yxl5lOeGPLmJ/D9fLNJnixS//zgKLZySaaF0PY8RRLaTi2Pntz9sfGclPQOAqWvXTM1gSk0LRpFfpxrpt3L3IHEJ97LY7gTIi9DoM5I40Wrutc1PXx+eUksFyhYJ2dJbnHnUqcaDSOxETifXb4LPim0ZTnoqfm5+SDUTxOvIHEVHWL3W6a3v4H7XogbPtdxsYQrPSR0EfIZ5QkG0kXW8C/+Ph9NNBx0Ve9AtYa01MYMf1pkOvr5PhE2XVfe3wU7C/XwQucfS+qDNlnnv6GnXWu3tI4fb69G57sSHmyglm9V0RCOIV9yYoA+LQssYRxCCZY8msJUt+jso9Rgr8f2GLYSLfiOq+zJzB+9IPrKb0+wHVrZEfiI1HrLpzgGs+yK1nJ2C0GWJSRw==]]></Encrypt>
</xml>`),
	}
	return &res, nil
}

var addr = flag.String("addr", "0.0.0.0:80", "server port")

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	pb.RegisterGreeterServer(s, &server{})
	pbLibrary.RegisterLibraryServer(s, &grpc_showcase.LibraryServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
