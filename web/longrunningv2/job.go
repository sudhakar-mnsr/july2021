package models

import (
	"time"

	"github.com/google/uuid"
)

// Job represents UUID of a Job
type Job struct {
	ID        uuid.UUID   `json:"uuid"`
	Type      string      `json:"type"`
	ExtraData interface{} `json:"extra_data"`
}

// Log is for Worker-A
type Log struct {
	ClientTime time.Time `json:"client_time"`
}

