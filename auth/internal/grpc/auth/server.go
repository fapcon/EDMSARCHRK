package auth

import (
	"auth/internal/service"
	"context"
	authpr "github.com/fapcon/MSHUGOprotos/protos/auth/gen"
	"log"
)

type ServicerAuth interface {
	Register(email, password string) (string, error)
	Login(email, password string) (string, error)
	ItsValid(token string) (bool, error)
}

type ServiceAuth struct {
	authpr.UnimplementedAuthServiceServer
	auths service.AuthService
}

func (s *ServiceAuth) Register(ctx context.Context, req *authpr.RegisterRequest) (*authpr.RegisterResponse, error) {
	mess, err := s.auths.Register(req.Email, req.Hashedpassword)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}

	return &authpr.RegisterResponse{Message: mess}, nil
}

func (s *ServiceAuth) Login(ctx context.Context, req *authpr.LoginRequest) (*authpr.LoginResponse, error) {
	token, err := s.auths.Login(req.Email, req.Password)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return &authpr.LoginResponse{Token: token}, nil
}

func (s *ServiceAuth) IsValid(ctx context.Context, req *authpr.ValidRequest) (*authpr.ValidResponse, error) {
	isvalid, err := s.auths.IsValid(req.Token)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return &authpr.ValidResponse{IsValid: isvalid}, nil
}
