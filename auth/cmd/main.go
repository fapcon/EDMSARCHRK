package main

import (
	cnt "auth/internal/controller"
	"auth/internal/grpc/auth"
	"auth/internal/grpc/grpcclients"
	"auth/internal/router"
	"auth/internal/service"
	"fmt"
	authpr "github.com/fapcon/MSHUGOprotos/protos/auth/gen"
	"github.com/go-chi/chi"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
)

func main() {

	uscl := grpcclients.NewClientUser()
	authservice := service.NewAuthService(uscl)
	authcnt := cnt.NewHandleAuth(authservice)

	r := router.Route(authcnt)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(r *chi.Mux) {
		fmt.Println("Запуск сервера auth")
		http.ListenAndServe(":8082", r)
		defer wg.Done()
	}(r)

	go func() {
		listen, err := net.Listen("tcp", ":44972")
		if err != nil {
			log.Fatalf("Ошибка при прослушивании порта: %v", err)
		}

		server := grpc.NewServer()
		authpr.RegisterAuthServiceServer(server, &auth.ServiceAuth{})

		log.Println("Запуск gRPC сервера auth...")
		if err := server.Serve(listen); err != nil {
			log.Fatalf("Ошибка при запуске сервера: %v", err)
		}
		defer wg.Done()
	}()
	wg.Wait()
}
