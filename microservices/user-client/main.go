package main

import (
	"context"
	"log"
	"os"

	pb "github.com/andreylm/learning/microservices/user-service/proto/user"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

func main() {
	cmd.Init()

	client := pb.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)

	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "Your full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "company",
				Usage: "Your company",
			},
		),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {

			r, err := client.Create(context.TODO(), &pb.User{
				Name:     c.String("name"),
				Email:    c.String("email"),
				Password: c.String("password"),
				Company:  c.String("company"),
			})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Created: %s", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}

			for _, v := range getAll.Users {
				log.Println(v)
			}

			os.Exit(0)
		}),
	)

	if err := service.Run(); err != nil {
		log.Panicln(err)
	}
}
