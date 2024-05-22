package grpccon

import (
	"log"

	userPb "auth.services/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcClient(address string) (userPb.UserGrpcServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		return nil, err
	}

	client := userPb.NewUserGrpcServiceClient(conn)

	return client, nil
}
