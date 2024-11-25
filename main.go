package main

import (
	"log"
	"payment-api/routes"
	"net/http"
)

func main() {
    routes.RegisterRoutes()
    log.Println("Server running on port 8082")
    http.ListenAndServe(":8082", nil)
}