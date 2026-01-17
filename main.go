package main

import (
	"log"
	"net/http"
	"web-go-native/config"
)

func main() {
	config.ConnectDB()

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
