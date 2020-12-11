package main

import (
	"context"
	"log"

	proto "github.com/Hands-On-Restful-Web-services-with-Go/chapter11/asyncClient/proto"
	micro "github.com/micro/go-micro"
)

// ProcessEvent processes a weather alert
func ProcessEvent(ctx context.Context, event *proto.Event) error {
	log.Println("Got alert:", event)
	return nil
}
