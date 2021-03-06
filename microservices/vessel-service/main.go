package main

import (
	"fmt"
	"log"

	micro "github.com/micro/go-micro"

	"os"

	pb "github.com/andreylm/learning/microservices/vessel-service/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &handler{session})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
