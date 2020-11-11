package main_test

import (
	"testing"
	"net/http"
)

func TestGetOriginalURL(t *testing.T) {
	// make a dummy reques
	response, err := http.Get("http://localhost:8000/v1/short/1")
 
