package main

import (
	"log"
	"net/http"
	"web-go-native/config"
	"web-go-native/controller/homecontroller"
)

func main() {
	config.ConnectDB()

	// 1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
