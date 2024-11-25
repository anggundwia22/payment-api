package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"payment-api/models"
	"payment-api/services"
)

func Payment(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		response := ApiResponse{
			Status:  "error",
			Message: "Authorization header is missing",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrNotSupported
		}
	
		return []byte("your_secret_key"), nil
	})

	if err != nil || !token.Valid {
		response := ApiResponse{
			Status:  "error",
			Message: "Invalid token",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	var paymentRequest models.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&paymentRequest); err != nil {
		response := ApiResponse{
			Status:  "error",
			Message: "Invalid request body",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if paymentRequest.MerchantID == "" {
		response := ApiResponse{
			Status:  "error",
			Message: "MerchantID is required",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if paymentRequest.Amount <= 0 {
		response := ApiResponse{
			Status:  "error",
			Message: "Amount must be greater than 0",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}


	message, err := services.Payment(paymentRequest.MerchantID, paymentRequest.Amount)
	if err != nil {
		response := ApiResponse{
			Status:  "error",
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}


	response := ApiResponse{
		Status:  "success",
		Message: "Payment successful",
		Data:    message, 
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}