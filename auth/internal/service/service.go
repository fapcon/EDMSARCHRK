package service

import (
	"auth/internal/grpc/grpcclients"
	"context"
	userpr "github.com/fapcon/MSHUGOprotos/protos/user/gen"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

const SecretKey = "secretkey"

type AuthService struct {
	clientUser *grpcclients.ClientUser
}

func NewAuthService(grcl *grpcclients.ClientUser) *AuthService {
	return &AuthService{clientUser: grcl}
}

func (a *AuthService) Register(email, hashedpassword string) (string, error) {

	mess, err := a.clientUser.CallCreate(context.Background(), &userpr.CreateRequest{
		Email:          email,
		Hashedpassword: hashedpassword,
	})
	if err != nil {
		log.Fatal("err call user")
	}

	return mess.Message, nil
}

func (a *AuthService) Login(email, password string) (string, error) {

	_, err := a.clientUser.CallCheck(context.Background(), &userpr.CheckRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Println("err: ", err)
		return "", err
	}
	token, err := GenerateToken(email)
	if err != nil {
		log.Println("err generate token:", err)
		return "", err
	}

	return token, nil
}

func (a *AuthService) IsValid(token string) (bool, error) {
	return false, nil
}

func GenerateToken(email string) (string, error) {

	claims := &jwt.StandardClaims{
		Subject:   email,
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
