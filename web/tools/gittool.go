package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/levigross/grequests"
	"github.com/urfave/cli"
)

var GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
var requestOptions = &grequests.RequestOptions{Auth: []string{GITHUB_TOKEN, "x-oauth-basic"}}

// Struct for holding response of repositories fetch API
type Repo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Forks    int    `json:"forks"`
	Private  bool   `json:"private"`
}

// Structs for modelling JSON body in create Gist
type File struct {
	Content string `json:"content"`
}
