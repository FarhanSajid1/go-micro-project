package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"

	user "github.com/farhansajid1/go-micro-project/user-service/proto/user"
	k8s "github.com/micro/examples/kubernetes/go/micro"

	"github.com/micro/go-micro"
)

// this package is for being the listener for the pubsub and "creating an email"
var topic = "user.created"

// struct for implementing subscriber
// open interface meaning we can implement any methods..
type subscriber struct{}

func (s *subscriber) Process(ctx context.Context, req *user.User) error {
	log.Println("Picked up a new message")
	log.Printf("sending out an email to %v", req.Email)
	return nil
}

var K8S = getEnv("K8S", "false")

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

type Service interface {
	Init(...micro.Option)
	Options() micro.Options
	Client() client.Client
	Server() server.Server
	Run() error
	String() string
}

func main() {
	// NewService() registers the service with the service discovery
	var service Service
	if K8S == "false" {
		service = micro.NewService(
			micro.Name("go.micro.srv.user"),
			micro.RegisterTTL(time.Second*30),
			micro.RegisterInterval(time.Second*10),
		)
	} else {
		service = k8s.NewService(
			micro.Name("go.micro.srv.user"),
			micro.RegisterTTL(time.Second*30),
			micro.RegisterInterval(time.Second*10),
		)
	}
	service.Init()

	micro.RegisterSubscriber(topic, service.Server(), &subscriber{})
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
