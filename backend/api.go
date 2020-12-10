package main

import (
	"log"
	"net/http"
)

type RequestPayload struct {
	API  string      `json:"api"`
	Data interface{} `json:"data"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

type GetNodeInfoPayload struct{}

func getNodeInfo(w http.ResponseWriter, r *http.Request) {}
