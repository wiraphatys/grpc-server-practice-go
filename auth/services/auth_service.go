package services

import (
	"context"
	"fmt"

	"auth.services/config"
	"auth.services/entities"
	userPb "auth.services/proto"
	"auth.services/repositories"
)

type authService struct {
	authRepo repositories.AuthRepository
	cfg      *config.Config
}

func NewAuthService(authRepo repositories.AuthRepository, cfg *config.Config) AuthService {
	return &authService{
		authRepo: authRepo,
		cfg:      cfg,
	}
}

func (s *authService) VerifyUser(pctx context.Context, req *entities.LoginData) (*entities.User, error) {
	reqRPC := &userPb.FindUserRequest{
		Email: req.Email,
	}
	existedUser, err := s.authRepo.FindUserByEmailGRPC(pctx, s.cfg.Grpc.UserUrl, reqRPC)
	if err != nil {
		return nil, err
	}

	if existedUser.Password == req.Password {
		return existedUser, nil
	}

	return nil, fmt.Errorf("error: %v", "invalid credentials")
}
