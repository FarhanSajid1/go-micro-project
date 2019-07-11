package main

import (
	"github.com/farhansajid1/go-micro-project/consignment-service/database"
	"log"
	"time"

	"gopkg.in/mgo.v2"

	pb "github.com/farhansajid1/go-micro-project/consignment-service/proto"
	vessel "github.com/farhansajid1/go-micro-project/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
)

type server struct{}

var collection *mgo.Collection
var err error

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// database
	log.Printf("Connecting to the database")
	session := database.ConnectDB()
	defer session.Close()
	collection = session.DB("consignments").C("new")
	log.Printf("connected successfully")

	service.Init()
	pb.RegisterShippingServiceHandler(service.Server(), new(server))

	// Register the client for the vessel.
	vesselClient = vessel.NewVesselService("go.micro.srv.vessel", service.Client())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}

