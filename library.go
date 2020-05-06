package grpc_showcase

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"log"

	pb "github.com/longkai/grpc-showcase/genproto/apis/library/v1"
)

type LibraryServer struct {
	pb.UnimplementedLibraryServer
}

func (l *LibraryServer) CreateShelf(ctx context.Context, req *pb.CreateShelfRequest) (*pb.Shelf, error) {
	return req.Shelf, nil
}

func (l *LibraryServer) GetShelf(ctx context.Context, req *pb.GetShelfRequest) (*pb.Shelf, error) {
	return &pb.Shelf{
		Name:  req.Name,
		Theme: "love feelings",
	}, nil
}

func (l *LibraryServer) ListShelves(ctx context.Context, req *pb.ListShelvesRequest) (*pb.ListShelvesResponse, error) {
	li := []*pb.Shelf{
		&pb.Shelf{
			Name:  "computation science",
			Theme: "theme computation science",
		},
		&pb.Shelf{
			Name:  "love feelings",
			Theme: "theme love feelings",
		},
	}
	return &pb.ListShelvesResponse{
		Shelves:       li,
		NextPageToken: "awesome next page token",
	}, nil
}

func (l *LibraryServer) DeleteShelf(ctx context.Context, req *pb.DeleteShelfRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (l *LibraryServer) MergeShelves(ctx context.Context, req *pb.MergeShelvesRequest) (*pb.Shelf, error) {
	return &pb.Shelf{
		Name:  req.OtherShelfName,
		Theme: "theme moved shelf",
	}, nil
}

func (l *LibraryServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.Book, error) {
	return req.Book, nil
}

func (l *LibraryServer) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.Book, error) {
	return &pb.Book{
		Name:   req.Name,
		Author: "Martin Kleppmann",
		Title:  "Designing Data-Intensive Applications",
		Read:   true,
	}, nil
}

func (l *LibraryServer) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	return &pb.ListBooksResponse{
		Books:         []*pb.Book{
			&pb.Book{
				Name:   "shelves/cs/books/ddia",
				Author: "Martin Kleppmann",
				Title:  "Designing Data-Intensive Applications",
				Read:   true,
			},
			&pb.Book{
				Name:   "shelves/cs/books/tcpip",
				Author: "W. Richard Stevens",
				Title:  "TCP/IP Illustrated, Vol. 1: The Protocols",
				Read:   true,
			},
		},
		NextPageToken: "awesome next page token",
	}, nil
}

func (l *LibraryServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (l *LibraryServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.Book, error) {
	// TODO: patch it
	log.Printf("UpdateBook: %+v", req)
	return req.Book, nil
}

func (l *LibraryServer) MoveBook(ctx context.Context, req *pb.MoveBookRequest) (*pb.Book, error) {
	return &pb.Book{
		Name:   req.Name,
		Author: "W. Richard Stevens",
		Title:  "TCP/IP Illustrated, Vol. 1: The Protocols",
		Read:   true,
	}, nil
}
