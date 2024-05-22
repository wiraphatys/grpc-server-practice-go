package repositories

import (
	"context"

	"auth.services/entities"
	userPb "auth.services/proto"
)

type AuthRepository interface {
	FindUserByEmailGRPC(pctx context.Context, grpcUrl string, req *userPb.FindUserRequest) (*entities.User, error)
}
