package controllers

import (
	"encoding/json"
	"net/http"

	"payment-api/models"
	"payment-api/services"
)

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` 
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		response := ApiResponse{
			Status:  "error",
			Message: "Invalid request body",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if loginRequest.Username == "" {
		response := ApiResponse{
			Status:  "error",
			Message: "Username is required",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if loginRequest.Password == "" {
		response := ApiResponse{
			Status:  "error",
			Message: "Password is required",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	customer, err := services.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		response := ApiResponse{
			Status:  "error",
			Message: "Username or Password is incorrect",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}


	token, err := services.GenerateToken(customer.ID) 
	if err != nil {
		response := ApiResponse{
			Status:  "error",
			Message: "Unable to generate token",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := ApiResponse{
		Status:  "success",
		Message: "Login successful",
		Data: map[string]interface{}{
			"username": customer.Username,
			"token":    token,
		},
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if err := services.Logout(); err != nil {
		response := ApiResponse{
			Status:  "error",
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := ApiResponse{
		Status:  "success",
		Message: "Logged out successfully",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}