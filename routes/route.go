package routes

import (
	"payment-api/controllers"
	"net/http"
)

func RegisterRoutes() {
    http.HandleFunc("/login", controllers.Login)
    http.HandleFunc("/logout", controllers.Logout)
    http.HandleFunc("/payment", controllers.Payment) 
}