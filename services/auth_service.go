package services

import (
	"errors"

	"payment-api/models"
	"payment-api/repositories"
)

var loggedInCustomer *models.Customer

func Login(username, password string) (*models.Customer, error) {
    customers, err := repositories.ReadCustomers()
    if err != nil {
        return nil, err
    }

    for _, customer := range customers {
        if customer.Username == username && customer.Password == password {
            loggedInCustomer = &customer
            return &customer, nil
        }
    }
    return nil, errors.New("invalid credentials")
}


func GetLoggedInCustomer() *models.Customer {
    return loggedInCustomer
}
func Logout() error {
    loggedInCustomer = nil
    return nil 
}