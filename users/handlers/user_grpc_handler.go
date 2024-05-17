package handlers

import (
	"context"

	userPb "user.services/proto"
	"user.services/services"
)

type userGrpcHandler struct {
	userService services.UserService
	userPb.UnimplementedUserGrpcServiceServer
}

func NewUserGrpcHandler(userService services.UserService) *userGrpcHandler {
	return &userGrpcHandler{
		userService: userService,
	}
}

func (g *userGrpcHandler) GetUserByEmailGRPC(ctx context.Context, req *userPb.FindUserRequest) (*userPb.UserProfile, error) {
	result, err := g.userService.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	return &userPb.UserProfile{
		Id:    int32(result.UID),
		Email: result.Email,
		Tel:   result.Tel,
	}, nil
}
