package main

import (
	"context"
	"log"

	user "github.com/farhansajid1/go-micro-project/user-service/proto/user"

	"github.com/micro/go-micro"
)

// this package is for being the listener for the pubsub and "creating an email"
var topic = "user.created"

// struct for implementing subscriber
// open interface meaning we can implement any methods..
type subscriber struct{}

func (s *subscriber) Process(ctx context.Context, req *user.User) error {
	log.Println("Picked up a new message")
	log.Println("sending out an email to %v", req.Email)
	return nil
}

func main() {
	// NewService() registers the service with the service discovery
	service := micro.NewService(
		micro.Name("go.micro.srv.email"),
	)
	service.Init()

	micro.RegisterSubscriber(topic, service.Server(), &subscriber{})
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
