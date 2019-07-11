package main

import (
	"context"
	"fmt"

	pb "github.com/farhansajid1/go-micro-project/consignment-service/proto"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("consignment.client"))
	service.Init()

	// name of the server is "greeter"
	client := pb.NewShippingService("go.micro.srv.consignment", service.Client())

	create := &pb.Consignment{
		Id:          "1",
		Description: "My first container",
		Weight:      120,
		Containers: []*pb.Container{
			&pb.Container{
				Id:         "1",
				CustomerId: "1",
				Origin:     "Pakistan",
				UserId:     "1",
			},
		},
		VesselId: "1",
	}

	resp, err := client.CreateConsignment(context.Background(), create)
	if err != nil {
		fmt.Printf("The error received was \n%v", err)
	}

	newresp, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		fmt.Printf("couldn't get all the items \n%v", err)
	}

	fmt.Printf("==> %v\t%v\t%v", resp.Created, resp.Consignment.VesselId, resp.Consignment.Id)
	fmt.Printf("==> %v\t", newresp.Consignments)

}
