package main

import (
	pb "github.com/andreylm/learning/microservices/consignment-service/proto/consignment"
	mgo "gopkg.in/mgo.v2"
)

const (
	dbName                = "shippy"
	consignmentCollection = "consignments"
)

// Repository - interface for db
type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

// ConsignmentRepository - mongodb repository
type ConsignmentRepository struct {
	session *mgo.Session
}

func (repo *ConsignmentRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(consignmentCollection)
}

// Create - add new consignment
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error {
	return repo.collection().Insert(consignment)
}

// GetAll - gets all consignments
func (repo *ConsignmentRepository) GetAll() (consignments []*pb.Consignment, err error) {
	err = repo.collection().Find(nil).All(&consignments)
	return
}

// Close - close sessions
func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}
