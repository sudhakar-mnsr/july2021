package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// Create new app
	app := cli.NewApp()

	// add flags with three arguments
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Value: "stranger",
			Usage: "your wonderful name",
		},
