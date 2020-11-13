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

// GetPackage fetches a package
func (driver *DBClient) GetPackage(w http.ResponseWriter, r *http.Request) {
	var Package = helper.Package{}
	vars := mux.Vars(r)
	// Handle response details
	driver.db.First(&Package, vars["id"])
	var PackageData interface{}
	// Unmarshal JSON string to interface
	json.Unmarshal([]byte(Package.Data), &PackageData)
	var response = PackageResponse{Package: Package}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(response)
	w.Write(respJSON)
}
