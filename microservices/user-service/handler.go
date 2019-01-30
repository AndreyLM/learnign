package main

import (
	pb "github.com/andreylm/learning/microservices/user-service/proto/user"
	"golang.org/x/net/context"
)

// Authable - interface for auth
type Authable interface {
}

// Service - implementation of proto methods
type Service struct {
	repo         Repository
	tokenService Authable
}

// Get - gets user
func (s *Service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := s.repo.Get(req.Id)
	if err != nil {
		return err
	}

	res.User = user
	return nil
}

// GetAll - gets all users
func (s *Service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	res.Users = users
	return nil
}

// Auth - authentication
func (s *Service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	// user, err := s.repo.GetByEmailAndPassword(req)
	_, err := s.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}

	res.Token = "test_token"
	return nil
}

// Create - creates new user
func (s *Service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.User = user
	return nil
}

// ValidateToken - validates token
func (s *Service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
