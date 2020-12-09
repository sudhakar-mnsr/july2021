package main

import (
	"context"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

// Response of APi
type Response struct {
	License struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"license"`
}
