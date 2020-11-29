package main

import (
	"log"
	"time"

	"encoding/json"

	"github.com/Hands-On-Restful-Web-services-with-Go/chapter9/longRunningTaskV1/models"
	"github.com/streadway/amqp"
)

// Workers does the job. It holds a connection
type Workers struct {
	conn *amqp.Connection
}

func (w *Workers) run() {
	log.Printf("Workers are booted up and running")
	channel, err := w.conn.Channel()
	handleError(err, "Fetching channel failed")
	defer channel.Close()
