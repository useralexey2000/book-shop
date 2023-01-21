package main

import (
	"book-shop/infra"
	"book-shop/infra/db"
	"book-shop/proto/pb"
	"fmt"
	"net"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// /docker-entrypoint-initdb.d

var (
	// host             = "0.0.0.0"
	host             = ""
	port             = os.Getenv("BOOKSERV_SVC_SERVICE_PORT")
	postgresHost     = os.Getenv("POSTGRES_HOST")
	postgresPort     = os.Getenv("POSTGRES_PORT")
	postgresDBName   = os.Getenv("POSTGRES_DB")
	postgresUser     = os.Getenv("POSTGRES_USER")
	postgresPassword = os.Getenv("POSTGRES_PASSWORD")
)

func main() {

	// local deployment
	// var (
	// host             = "localhost"
	// port             = "9000"
	// postgresHost     = "localhost"
	// postgresPort     = "5432"
	// postgresDBName   = "mydb"
	// postgresUser     = "pguser"
	// postgresPassword = "pgsecret"
	// )

	fmt.Printf("host an port are: %v, %v\n", host, port)

	cfg, err := pgxpool.ParseConfig(
		fmt.Sprintf(
			"postgres://%v:%v@%v:%v/%v?sslmode=disable&pool_max_conns=1",
			postgresUser, postgresPassword, postgresHost, postgresPort, postgresDBName))

	if err != nil {
		panic(fmt.Sprintf("cant parse conf, %v", err))
	}

	fmt.Println("db config parsed", cfg.ConnString())

	repo, err := db.NewPostgresRepo(cfg)
	if err != nil {
		panic(fmt.Sprintf("cant init repo, %v", err))
	}

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

	fmt.Printf("starting books service, port: ß%v, host: %v", port, host)
	if err = s.Serve(ls); err != nil {
		panic(fmt.Sprintf("cant serve grpc, %v", err))
	}
}
