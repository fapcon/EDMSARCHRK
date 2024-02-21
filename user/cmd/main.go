package main

import (
	"fmt"
	userpr "github.com/fapcon/MSHUGOprotos/protos/user/gen"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
	cnt "user/internal/controller"
	"user/internal/grpc/user"
	"user/internal/repository"
	"user/internal/router"
	"user/internal/service"
)

func main() {
	time.Sleep(2 * time.Second)
	dbHost := "db"
	dbPort := "5432"
	dbUser := "userpostgres"
	dbPassword := "password"
	dbName := "userserv"
	sslmode := "disable"

	connectionString := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=" + sslmode

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	if err != nil {
		log.Fatalf("ping:%v", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, email VARCHAR(255) UNIQUE NOT NULL, hashedpassword VARCHAR(255) NOT NULL)`)

	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepo(db)

	userService := service.NewUserService(userRepo)

	serviceUser := user.NewServiceUser(userService)

	usercnt := cnt.NewHandleUser(userService)

	r := router.Route(usercnt)

	w := sync.WaitGroup{}
	w.Add(2)
	go func(r *chi.Mux) {
		fmt.Println("Запуск user сервера 8083")
		http.ListenAndServe(":8083", r)
		defer w.Done()
	}(r)

	go func() {
		lis, err := net.Listen("tcp", ":44971")
		if err != nil {
			log.Fatalf("Failed to listen on port %s: %v", "44971", err)
		}
		grpcServer := grpc.NewServer()

		userpr.RegisterUserServiceServer(grpcServer, serviceUser)

		log.Print("Starting gRPC server user...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
		defer w.Done()
	}()
	w.Wait()
}
