package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"grpc-vimn-demo-go/library"
	pb "grpc-vimn-demo-go/protos"
)

const SERVER_ADDRESS = "localhost:8092"

type libraryServer struct {
	pb.UnimplementedLibraryServer
}

func (s libraryServer) ListBooks(request *pb.ListBooksRequest, server pb.Library_ListBooksServer) error {
	bookQueryBuilder := library.NewBookQueryBuilder()

	books := bookQueryBuilder.GetBooks()
	for _, book := range books {
		server.Send(book)
	}
	return nil
}

func (s libraryServer) FilterBooks(request *pb.FilterBooksRequest, server pb.Library_FilterBooksServer) error {
	bookQueryBuilder := library.NewBookQueryBuilder()
	bookQueryBuilder.FilterTitle(request.Title.GetValue())
	bookQueryBuilder.FilterAuthorName(request.AuthorName.GetValue())

	books := bookQueryBuilder.GetBooks()
	for _, book := range books {
		server.Send(book)
	}
	return nil
}

func (s libraryServer) SubscribeForBookUpdates(request *pb.SubscribeForBookUpdatesRequest, server pb.Library_SubscribeForBookUpdatesServer) error {
	c := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 5)
		server.Send(&pb.Book{Id: 100, Title: "New Book"})
		close(c)
	}()

	<-c
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
