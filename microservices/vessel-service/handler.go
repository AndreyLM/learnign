package main

import (
	"golang.org/x/net/context"

	pb "github.com/andreylm/learning/microservices/vessel-service/proto/vessel"
	mgo "gopkg.in/mgo.v2"
)

type handler struct {
	session *mgo.Session
}

func (s *handler) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	vessel, err := repo.FindAvailable(req)

	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func (s *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	if err := repo.Create(req); err != nil {
		return err
	}

	res.Created = true
	res.Vessel = req
	return nil
}
