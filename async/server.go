package main

import (
	"context"
	"log"
	"time"

	proto "github.com/Hands-On-Restful-Web-services-with-Go/chapter11/asyncService/proto"
	micro "github.com/micro/go-micro"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("weather"),
	)
	p := micro.NewPublisher("alerts", service.Client())

