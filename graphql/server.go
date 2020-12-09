package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Player holds player response
type Player struct {
	ID             int      `json:"int"`
	Name           string   `json:"name"`
	HighScore      int      `json:"highScore"`
	IsOnline       bool     `json:"isOnline"`
	Location       string   `json:"location"`
	LevelsUnlocked []string `json:"levelsUnlocked"`
}

var players = []Player{
	Player{ID: 123, Name: "Pablo", HighScore: 1100, IsOnline: true, Location: "Italy"},
	Player{ID: 230, Name: "Dora", HighScore: 2100, IsOnline: false, Location: "Germany"},
}

var playerObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
