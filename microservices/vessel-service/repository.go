package main

import (
	pb "github.com/andreylm/learning/microservices/vessel-service/proto/vessel"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

// Repository - interface for db
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

// VesselRepository - mongodb repository
type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}

// Create - add new consignment
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

// FindAvailable - gets all consignments
func (repo *VesselRepository) FindAvailable(specification *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel
	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": specification.Capacity},
		"maxweight": bson.M{"$gte": specification.MaxWeight},
	}).One(&vessel)
	if err != nil {
		return nil, err
	}

	return vessel, nil
}

// Close - close sessions
func (repo *VesselRepository) Close() {
	repo.session.Close()
}
