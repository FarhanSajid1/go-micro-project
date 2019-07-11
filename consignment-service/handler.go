package main

import (
	"context"
	"log"

	"gopkg.in/mgo.v2/bson"

	pb "github.com/farhansajid1/go-micro-project/consignment-service/proto"
	vessel "github.com/farhansajid1/go-micro-project/vessel-service/proto/vessel"
)

type ConsignmentItem struct {
	// primitive.ObjectID is the id of the bson object
	ID          bson.ObjectId   `bson:"_id,omitempty"` // generated on the fly
	Description string          `bson:"description"`
	Weight      int32           `bson:"weight"`
	VesselId    string          `bson:"vessel_id"`
	Containers  []*pb.Container `bson:"container"`
}

/*string id = 1;
  string description = 2;
  int32 weight = 3;
  repeated Container containers = 4;
  string vessel_id = 5; */

var vesselClient vessel.VesselService

func createVesselReq() *vessel.Specification {
	ves := &vessel.Specification{
		MaxWeight: 100,
		Capacity:  500,
	}
	return ves
}

func toBSON(req *pb.Consignment, item *pb.Container) *ConsignmentItem {
	// convert req to bson object
	i := bson.NewObjectId()
	obj := &ConsignmentItem{
		ID:          i,
		Description: req.GetDescription(),
		Weight:      req.GetWeight(),
		Containers: []*pb.Container{
			&pb.Container{
				Id:         item.GetId(),
				CustomerId: item.GetCustomerId(),
				Origin:     item.GetOrigin(),
				UserId:     item.GetUserId(),
			},
		},
	}
	return obj
}

func (s *server) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	lis := req.GetContainers()
	item := lis[0]

	// Getting vessel id
	ves := createVesselReq()
	vesselresp, err := vesselClient.FindAvailable(ctx, ves)
	if err != nil {
		log.Printf("there was an error %v", err)
	}

	log.Printf("Found a vessel %v\t%v", vesselresp.Vessel.Descriptor, vesselresp.Vessel.Id)

	obj := toBSON(req, item)
	err = collection.Insert(obj)
	if err != nil{
		log.Print("error writing!!!! %v", err)
	}

	create := &pb.Consignment{
		Id:          obj.ID.String(),
		Description: req.GetDescription(),
		Weight:      req.GetWeight(),
		Containers: []*pb.Container{
			&pb.Container{
				Id:         item.GetId(),
				CustomerId: item.GetCustomerId(),
				Origin:     item.GetOrigin(),
				UserId:     item.GetUserId(),
			},
		},
		VesselId: vesselresp.Vessel.Id,
	}
	// b, err := proto.Marshal(create)
	// if err != nil {
	// 	log.Printf("could not write to bytes %v", err)
	// }
	// ioutil.WriteFile("./input", b, 0644)

	resp.Created = true
	resp.Consignment = create
	return nil
}

func (s *server) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	// Returns all the blogs that are found
	var coll []*pb.Consignment
	err := collection.Find(nil).Limit(10).All(&coll)
	if err != nil {
		log.Printf("could not find anything %v", err)
	}

	resp.Consignments = coll

	return nil
}
