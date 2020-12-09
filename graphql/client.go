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

func main() {
	// create a client (safe to share across requests)
	client := graphql.NewClient("https://api.github.com/graphql")

	// make a request to GitHub API
	req := graphql.NewRequest(`
		query {
			license(key: "apache-2.0") {
				name
				description
			}
		}
`)
