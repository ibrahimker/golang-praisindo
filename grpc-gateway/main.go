package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ibrahimker/golang-praisindo/grpc-gateway/common/config"
	"github.com/ibrahimker/golang-praisindo/grpc-gateway/common/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServer struct {
	model.UnimplementedUserServiceServer
}

var Users = make(map[string]*model.User)

func main() {
	srv := grpc.NewServer()
	userSrv := new(UserServer)
	model.RegisterUserServiceServer(srv, userSrv)

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

func (t *UserServer) GetAll(ctx context.Context, req *emptypb.Empty) (*model.GetAllResponse, error) {
	var todos []*model.User
	for _, v := range Users {
		todos = append(todos, &model.User{
			Username: v.GetUsername(),
			Email:    v.GetEmail(),
			Password: v.GetPassword(),
			Age:      v.GetAge(),
		})
	}
	return &model.GetAllResponse{Data: todos}, nil
}
func (t *UserServer) GetByEmail(ctx context.Context, req *model.GetByEmailRequest) (*model.GetByEmailResponse, error) {
	todo, ok := Users[req.GetEmail()]
	if !ok {
		return nil, errors.New("not found")
	}
	return &model.GetByEmailResponse{Data: todo}, nil
}
func (t *UserServer) Create(ctx context.Context, req *model.User) (*model.MutationResponse, error) {
	Users[req.GetEmail()] = &model.User{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Age:      req.GetAge(),
	}
	msg := req.GetUsername() + "successfully appended"
	return &model.MutationResponse{Success: msg}, nil
}
func (t *UserServer) Update(ctx context.Context, req *model.UpdateRequest) (*model.MutationResponse, error) {
	// TODO: implement me
	// Users[req.GetEmail()] = &model.User{
	// 	Id:   req.Get(),
	// 	Name: req.GetName(),
	// }
	msg := req.GetEmail() + "successfully appended"
	return &model.MutationResponse{Success: msg}, nil
}
func (t *UserServer) Delete(ctx context.Context, req *model.DeleteRequest) (*model.MutationResponse, error) {
	delete(Users, req.GetEmail())
	msg := req.GetEmail() + "successfully deleted"
	return &model.MutationResponse{Success: msg}, nil
}
