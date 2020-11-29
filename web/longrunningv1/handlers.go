package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Hands-On-Restful-Web-services-with-Go/chapter9/longRunningTaskV1/models"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

// JobServer holds handler functions
type JobServer struct {
	Queue   amqp.Queue
	Channel *amqp.Channel
	Conn    *amqp.Connection
}
