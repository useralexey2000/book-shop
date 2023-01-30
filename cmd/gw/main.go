package main

import (
	"book-shop/gw"
	"book-shop/proto/pb"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

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
	fmt.Printf("env host %v:%v, grpc %v:%v\n: ", host, port, grpcHost, grpcPort)

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx, net.JoinHostPort(grpcHost, grpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithInsecure(),
		// grpc.WithBlock(),
	)
	// defer cancel()

	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	if err != nil {
		panic(fmt.Sprint("cant dial grpc server ", err))
	}

	var quit bool

	go func() {
		<-time.After(20 * time.Second)
		quit = true
	}()
	go func() {
		for !quit {
			time.Sleep(1 * time.Second)
			fmt.Printf("grpc conn state : %v\n", conn.GetState().String())
		}
	}()

	client := pb.NewBookServiceClient(conn)
	gateway := gw.New(client)

	fmt.Println("client created")
	http.HandleFunc("/books/list", gateway.ListBooks())
	http.HandleFunc("/books/create", gateway.CreateBook())

	fmt.Printf("gw server is started on address: %v:%v\n", host, port)
	err = http.ListenAndServe(net.JoinHostPort(host, port), nil)
	if err != nil {
		fmt.Println("gw server closed with err: ", err)
	}
}
