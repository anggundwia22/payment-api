package repositories

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"payment-api/models"
)

func ReadCustomers() ([]models.Customer, error) {
    file, err := os.Open("data/customers.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    byteValue, _ := ioutil.ReadAll(file)

    var customers []models.Customer
    json.Unmarshal(byteValue, &customers)
    return customers, nil
}

func WriteCustomers(customers []models.Customer) error {
    data, err := json.MarshalIndent(customers, "", "  ")
    if err != nil {
        return err
    }
    err = ioutil.WriteFile("data/customers.json", data, 0644)
    return err
}

func UpdateCustomerBalance(updatedCustomer *models.Customer) error {
    customers, err := ReadCustomers()
    if err != nil {
        return err
    }

    for i, customer := range customers {
        if customer.ID == updatedCustomer.ID {
            customers[i].Balance = updatedCustomer.Balance
            break
        }
    }

    return WriteCustomers(customers)
}


func AddHistory(history models.History) error {
    var histories []models.History

    if _, err := os.Stat("data/history.json"); err == nil {
        file, err := os.Open("data/history.json")
        if err != nil {
            return err
        }
        defer file.Close()

        byteValue, _ := ioutil.ReadAll(file)
        json.Unmarshal(byteValue, &histories)
    }

    histories = append(histories, history)

    data, err := json.MarshalIndent(histories, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile("data/history.json", data, 0644)
}
