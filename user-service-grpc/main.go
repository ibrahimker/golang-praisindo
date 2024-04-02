package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ibrahimker/golang-praisindo/user-service-grpc/common/config"
	"github.com/ibrahimker/golang-praisindo/user-service-grpc/common/model"
	grpchandler "github.com/ibrahimker/golang-praisindo/user-service-grpc/handler/grpc"
	"github.com/ibrahimker/golang-praisindo/user-service-grpc/repository"
	"github.com/ibrahimker/golang-praisindo/user-service-grpc/service"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	// init database connection
	dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// setup repo-service-handler
	userRepository := repository.NewUserRepository(db)
	userServices := service.NewUserSvc(userRepository)
	// use grpc handler
	userHandler := grpchandler.NewUserGRPCHandler(userServices)

	srv := grpc.NewServer()
	model.RegisterUserServiceServer(srv, userHandler)

	log.Println("Starting User Server at ", config.ServiceUserPort)

	listener, err := net.Listen("tcp", config.ServiceUserPort)
	if err != nil {
		log.Fatalf("could not listen. Err: %+v\n", err)
	}

	// setup http proxy
	go func() {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		grpcServerEndpoint := flag.String("grpc-server-endpoint", "localhost"+config.ServiceUserPort, "gRPC server endpoint")
		_ = model.RegisterUserServiceHandlerFromEndpoint(context.Background(), mux, *grpcServerEndpoint, opts)
		log.Println("Starting User Server HTTP at 9001 ")
		http.ListenAndServe(":9001", mux)
	}()

	log.Fatalln(srv.Serve(listener))
}
