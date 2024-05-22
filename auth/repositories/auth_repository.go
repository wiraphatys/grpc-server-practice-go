package repositories

import (
	"context"
	"log"
	"time"

	"auth.services/config"
	"auth.services/entities"
	"auth.services/grpccon"
	userPb "auth.services/proto"
	"google.golang.org/grpc/metadata"
)

type authRepository struct {
	cfg *config.Config
}

func NewAuthRepository(cfg *config.Config) AuthRepository {
	return &authRepository{
		cfg: cfg,
	}
}

func (r *authRepository) FindUserByEmailGRPC(pctx context.Context, grpcUrl string, req *userPb.FindUserRequest) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("auth", r.cfg.Key.ApiKey))
	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %v", err.Error())
		return nil, err
	}

	result, err := conn.GetUserByEmailGRPC(ctx, req)
	if err != nil {
		log.Printf("Error: GetUserByEmailGRPC failed: %v", err.Error())
		return nil, err
	}

	return &entities.User{
		UID:      int(result.Id),
		Email:    result.Email,
		Password: result.Password,
		Tel:      result.Tel,
	}, nil
}
