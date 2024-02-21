package main

import (
	cnt "geo/internal/controller"
	"geo/internal/grpc/geo"
	"geo/internal/router"
	"geo/internal/service"
	geopr "github.com/fapcon/MSHUGOprotos/protos/geo/gen"
	"github.com/go-chi/chi"
	"github.com/streadway/amqp"
	"gitlab.com/ptflp/gopubsub/kafkamq"
	"gitlab.com/ptflp/gopubsub/queue"
	"gitlab.com/ptflp/gopubsub/rabbitmq"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

func main() {
	var MQ queue.MessageQueuer
	eventStreamer := os.Getenv("EVENTSTREAMER")
	switch eventStreamer {
	case "kafka":
		MQK, err := kafkamq.NewKafkaMQ("kafka:9092", "myGroup")
		if err != nil {
			log.Fatal(err)
		}
		MQ = MQK
	case "rabbit":
		conn, err := amqp.Dial("amqp://guest:guest@rabbit:5672/")
		if err != nil {
			log.Fatal(err)
		}

		MQR, err := rabbitmq.NewRabbitMQ(conn)
		if err != nil {
			log.Fatal(err)
		}

		err = rabbitmq.CreateExchange(conn, "ratelimiter", "topic")
		if err != nil {
			log.Fatal(err)
		}
		MQ = MQR
	}

	geoservice := service.NewGeoService()
	geohandle := cnt.NewHandleGeo(geoservice)
	r := router.Route(geohandle, MQ)

	w := sync.WaitGroup{}
	w.Add(2)

	go func(r *chi.Mux) {
		defer w.Done()
		http.ListenAndServe(":8081", r)
	}(r)

	go func() {
		defer w.Done()
		listen, err := net.Listen("tcp", ":44973")
		if err != nil {
			log.Fatalf("Ошибка при прослушивании порта: %v", err)
		}

		server := grpc.NewServer()
		geopr.RegisterGeoServiceServer(server, &geo.ServerGeo{})

		log.Println("Запуск gRPC сервера geo...")
		if err := server.Serve(listen); err != nil {
			log.Fatalf("Ошибка при запуске сервера: %v", err)
		}
	}()
	w.Wait()
}
