package main

import (
	"flashpage/app/config"
	"flashpage/app/routes"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()

	config.ConnectDB()

	r := routes.RegisterRoutes()

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}