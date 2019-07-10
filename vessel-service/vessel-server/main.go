package main

import (
	"github.com/micro/go-micro/errors"
	"context"
	"log"
	"time"

	vessel "github.com/farhansajid1/go-micro-project/vessel-service/proto/vessel"

	"github.com/micro/go-micro"
)

type server struct{}
/* 
string id = 1;
int32 capacity = 2;
int32 max_weight = 3;
string name = 4;
bool available = 5;
string owner_id = 6;

*/

func (s *server) FindAvailable(ctx context.Context, req *vessel.Specification, resp *vessel.Response) error {
	ves := &vessel.Vessel{
		Id : "new-vessel-786",
		Capacity: 200000,
		MaxWeight: 500000,
		Name: "Farhan's Vessel",
		Available: true,
		OwnerId: "786",
	}

	if req.Capacity < ves.Capacity && req.MaxWeight < ves.MaxWeight {
		resp.Vessel = ves 	
	} else {
		return errors.BadRequest("go.micro.srv.vessel", "No vessels available..")
	}
	return nil

}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()
	vessel.RegisterVesselServiceHandler(service.Server(), new(server))

	// start up the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
