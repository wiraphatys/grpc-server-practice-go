package services

import (
	"context"

	"auth.services/entities"
)

type AuthService interface {
	VerifyUser(pctx context.Context, req *entities.LoginData) (*entities.User, error)
}
