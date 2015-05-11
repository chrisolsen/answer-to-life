package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Println("Failed to create request", err)
		t.Fail()
		return
	}

	h, _ := mux.Handler(req)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		log.Println("Bad status code", rr.Code)
		t.Fail()
		return
	}

	fmt.Println(rr.Body)
}
