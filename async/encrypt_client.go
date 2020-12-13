package main

import (
	"context"
	"fmt"

	proto "github.com/Hands-On-Restful-Web-services-with-Go/chapter11/encryptClient/proto"
	micro "github.com/micro/go-micro"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("encrypter.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new encrypter client
	encrypter := proto.NewEncrypterService("encrypter", service.Client())
