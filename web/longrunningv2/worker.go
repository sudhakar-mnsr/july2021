package main

import (
	"log"
	"time"

	"encoding/json"

	"github.com/Hands-On-Restful-Web-services-with-Go/chapter9/longRunningTaskV2/models"
	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
)

// Workers does the job. It holds connections
type Workers struct {
	conn        *amqp.Connection
	redisClient *redis.Client
}

func (w *Workers) run() {
	log.Printf("Workers are booted up and running")
	channel, err := w.conn.Channel()
