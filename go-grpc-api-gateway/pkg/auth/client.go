package auth

import (
	"fmt"

	"github.com/sabilimaulana/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/sabilimaulana/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect to auth service:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
