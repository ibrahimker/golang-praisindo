package grpchandler

import (
	"context"

	"github.com/ibrahimker/golang-praisindo/user-service-grpc/common/model"
	"github.com/ibrahimker/golang-praisindo/user-service-grpc/entity"
	"github.com/ibrahimker/golang-praisindo/user-service-grpc/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserGRPCServiceHandler struct {
	model.UnimplementedUserServiceServer
	userServices service.IUserService
}

func NewUserGRPCHandler(userServices service.IUserService) *UserGRPCServiceHandler {
	return &UserGRPCServiceHandler{
		userServices: userServices,
	}
}

func (u *UserGRPCServiceHandler) GetAll(ctx context.Context, req *emptypb.Empty) (*model.GetAllResponse, error) {
	users, err := u.userServices.GetAll()
	if err != nil {
		return nil, err
	}

	// convert data type
	resData := []*model.User{}
	for _, u := range users {
		resData = append(resData, &model.User{
			Username: u.Username,
			Email:    u.Email,
			Password: u.Password,
			Age:      int32(u.Age),
		})
	}
	return &model.GetAllResponse{
		Data: resData,
	}, nil
}

func (u *UserGRPCServiceHandler) GetByEmail(ctx context.Context, req *model.GetByEmailRequest) (*model.GetByEmailResponse, error) {
	user, err := u.userServices.GetByEmail(req.GetEmail())
	if err != nil {
		return nil, err
	}
	return &model.GetByEmailResponse{
		Data: &model.User{
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
			Age:      int32(user.Age),
		},
	}, nil
}

func (u *UserGRPCServiceHandler) Create(ctx context.Context, req *model.User) (*model.MutationResponse, error) {
	// TODO: implement me
	_, err := u.userServices.Create(entity.User{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Age:      int(req.GetAge()),
	})
	if err != nil {
		return nil, err
	}

	return &model.MutationResponse{
		Success: "successfully create user",
	}, nil
}

func (u *UserGRPCServiceHandler) Update(ctx context.Context, req *model.UpdateRequest) (*model.MutationResponse, error) {
	_, err := u.userServices.Update(entity.User{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Age:      int(req.GetAge()),
	})
	if err != nil {
		return nil, err
	}

	return &model.MutationResponse{
		Success: "successfully update user",
	}, nil
}

func (u *UserGRPCServiceHandler) Delete(ctx context.Context, req *model.DeleteRequest) (*model.MutationResponse, error) {
	if err := u.userServices.Delete(req.GetEmail()); err != nil {
		return nil, err
	}

	return &model.MutationResponse{
		Success: "successfully delete user",
	}, nil
}
