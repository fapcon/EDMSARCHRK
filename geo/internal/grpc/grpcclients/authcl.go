package grpcclients

import (
	"context"
	authpr "github.com/fapcon/MSHUGOprotos/protos/auth/gen"
	"google.golang.org/grpc"
	"log"
)

type ClientAuth struct{}

func (c *ClientAuth) CallIsValid(ctx context.Context, req *authpr.ValidRequest) (*authpr.ValidResponse, error) {
	conn, err := grpc.Dial("auth:44972", grpc.WithInsecure())
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	client := authpr.NewAuthServiceClient(conn)
	res, err := client.IsValid(ctx, req)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return res, nil
}
