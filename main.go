package main

import (
	"log"
	"net/http"
	"web-go-native/config"
	"web-go-native/controller/categorycontroller"
	"web-go-native/controller/homecontroller"
)

func main() {
	config.ConnectDB()

	// 1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Category
	http.HandleFunc("/catgories", categorycontroller.Index)
	http.HandleFunc("/category/add", categorycontroller.Add)
	http.HandleFunc("/category/edit", categorycontroller.Edit)
	http.HandleFunc("/category/delete", categorycontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
