package grpccon

import (
	"context"
	"errors"
	"log"
	"net"

	"user.services/config"
	// Import protobuf for user service
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type grpcAuth struct {
	secretKey string
}

func (g *grpcAuth) unaryAuthorization(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("Error: Metadata not found")
		return nil, errors.New("error: metadata not found")
	}

	authHeader, ok := md["auth"]
	if !ok {
		log.Printf("Error: Metadata not found")
		return nil, errors.New("error: metadata not found")
	}

	if len(authHeader) == 0 {
		log.Printf("Error: Metadata not found")
		return nil, errors.New("error: metadata not found")
	}

	// ตรงนี้คุณสามารถตรวจสอบ JWT token ได้
	// claims, err := jwtauth.ParseToken(g.secretKey, string(authHeader[0]))
	// if err != nil {
	//     log.Printf("Error: Parse token failed: %s", err.Error())
	//     return nil, errors.New("error: token is invalid")
	// }
	// log.Printf("claims: %v", claims)

	return handler(ctx, req)
}

func NewGrpcServer(cfg *config.Jwt, host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)
	grpcAuth := &grpcAuth{
		secretKey: cfg.Secret,
	}
	opts = append(opts, grpc.UnaryInterceptor(grpcAuth.unaryAuthorization))
	grpcServer := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Error: Failed to listen: %v", err)
	}
	return grpcServer, lis
}
