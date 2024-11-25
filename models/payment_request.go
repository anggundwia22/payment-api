package models

type PaymentRequest struct {
    MerchantID string  `json:"merchant_id"`
    Amount     float64 `json:"amount"`
}