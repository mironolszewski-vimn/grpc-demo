package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"grpc-vimn-demo-go/library"
	pb "grpc-vimn-demo-go/protos"
)

const SERVER_ADDRESS = "localhost:8092"

type libraryServer struct {
	pb.UnimplementedLibraryServer
}

func (s libraryServer) ListBooks(request *pb.ListBooksRequest, server pb.Library_ListBooksServer) error {
	for _, book := range library.ListBooks() {
		server.Send(book)
	}
	return nil
}

func (s libraryServer) FilterBooks(request *pb.FilterBooksRequest, server pb.Library_FilterBooksServer) error {
	books := library.FilterBooks(request.Title.GetValue(), request.AuthorName.GetValue())
	for _, book := range books {
		server.Send(book)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", SERVER_ADDRESS)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterLibraryServer(grpcServer, &libraryServer{})
	fmt.Printf("Starting server listening at %s...", SERVER_ADDRESS)
	grpcServer.Serve(lis)
}
