package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Hands-On-Restful-Web-services-with-Go/chapter7/jsonstore/helper"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// DBClient stores the database session imformation. Needs to be initialized once
type DBClient struct {
	db *gorm.DB
}

// PackageResponse is the response to be send back for Package
type PackageResponse struct {
	Package helper.Package `json:"Package"`
}
