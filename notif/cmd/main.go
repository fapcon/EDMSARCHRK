package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"gitlab.com/ptflp/gopubsub/kafkamq"
	"gitlab.com/ptflp/gopubsub/queue"
	"gitlab.com/ptflp/gopubsub/rabbitmq"
	"log"
	"notif/internal/models"
	"notif/internal/service"
	"os"
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
	message, err := MQ.Subscribe("ratelimiter")
	if err != nil {
		log.Fatal(err)
	}
	data := &models.Data{}
	ns := service.NewNotifService()
	for m := range message {

		err = MQ.Ack(&m)
		if err != nil {
			log.Println("no data")
		}
		err = json.Unmarshal(m.Data, &data)
		if err != nil {
			log.Println("err unmarshal")
		}
		emailresp, _ := ns.SendViaEmail(data.Email)
		log.Println(emailresp)

		smsresp, _ := ns.SendViaSMS(data.Phone)
		log.Println(smsresp)
	}

}
