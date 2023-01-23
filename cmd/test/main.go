package main

import (
	"book-shop/infra"
	db "book-shop/infra/db/mock"
	"book-shop/proto/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	host = "localhost"
	port = "9000"
)

func main() {

	fmt.Printf("host an port are: %v, %v\n", host, port)

	repo := &db.MockDB{}
	fmt.Println("db initialized")

	ls, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		panic(fmt.Sprintf("cant start tcp listener, %v", err))
	}

	fmt.Println("listener created")

	s := grpc.NewServer()
	srv := infra.NewLibraryService(repo)
	pb.RegisterBookServiceServer(s, srv)

	reflection.Register(s)

	fmt.Println("grpc books service registered")

	fmt.Printf("starting books service, port: %v, host: %v", port, host)
	if err = s.Serve(ls); err != nil {
		panic(fmt.Sprintf("cant serve grpc, %v", err))
	}
}
