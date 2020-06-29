package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"example/handler"
	"example/subscriber"

	example "example/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.example"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.example", service.Server(), new(subscriber.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
