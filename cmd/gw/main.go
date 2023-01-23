package main

import (
	"book-shop/gw"
	"book-shop/proto/pb"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// host, port = "localhost", "8080"
	host     = os.Getenv("GW_HOST")
	port     = os.Getenv("BOOKSERV_SVC_SERVICE_PORT_BOOKSERVGW")
	grpcHost = os.Getenv("GRPC_HOST")
	grpcPort = os.Getenv("GRPC_PORT")
)

func main() {
	fmt.Println("initializing gw/client-side book")

	conn, err := grpc.DialContext(
		context.Background(), net.JoinHostPort(grpcHost, grpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(fmt.Sprint("cant dial grpc server ", err))
	}

	client := pb.NewBookServiceClient(conn)
	gateway := gw.New(client)

	http.HandleFunc("/books/list", gateway.ListBooks())
	http.HandleFunc("/books/create", gateway.CreateBook())

	fmt.Printf("gw server is started on address: %v:%v", host, port)
	err = http.ListenAndServe(net.JoinHostPort(host, port), nil)
	if err != nil {
		fmt.Println("gw server closed with err: ", err)
	}
}
