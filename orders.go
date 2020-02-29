// order
package prom

import (
	"fmt"
)

type Order struct {
	Id               int    `json:"id"`
	DateCreated      string `json:"date_created"`
	ClientFirstName  string `json:"client_first_name"`
	ClientSecondName string `json:"client_second_name"`
	ClientLastName   string `json:"client_last_name"`
	ClientNotes      string `json:"client_notes"`
	Products         []struct {
		Id          int     `json:"id"`
		ExternalId  string  `json:"external_id"`
		Image       string  `json:"image"`
		Quantity    float32 `json:"quantity"`
		Price       string  `json:"price"`
		Name        string  `json:"name"`
		TotalPrice  string  `json:"total_price"`
		MeasureUnit string  `json:"measure_unit"`
		Sku         string  `json:"sku"`
	}
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Price          string `json:"price"`
	DeliveryOption struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"delivery_option"`
	DeliveryAddress string `json:"delivery_address"`
	PaymentOption   struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	Status string `json:"status"`
	Source string `json:"source"`
}

type Orders struct {
	Orders []Order `json:"orders"`
	Error  string  `json:"error"`
}

func (acc *PromAccount) RequestOrders(params map[string]string) (orders []Order, err error) {
	var result Orders

	err = acc.client.Request("orders/list", params, &result)
	if err != nil {
		return nil, fmt.Errorf("Error when request orders: %s", result.Error)
	}

	return result.Orders, nil
}
