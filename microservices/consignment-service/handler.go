package main

import (
	"log"

	"golang.org/x/net/context"

	pb "github.com/andreylm/learning/microservices/consignment-service/proto/consignment"
	vesselProto "github.com/andreylm/learning/microservices/vessel-service/proto/vessel"
	mgo "gopkg.in/mgo.v2"
)

type handler struct {
	session      *mgo.Session
	vesselClient vesselProto.VesselServiceClient
}

func (s *handler) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	if err != nil {
		return err
	}

	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)

	req.VesselId = vesselResponse.Vessel.Id
	err = repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	consignments, err := repo.GetAll()

	if err != nil {
		return err
	}

	res.Consignments = consignments
	return nil
}
