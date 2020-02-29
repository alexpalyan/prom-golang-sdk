package prom

import (
	"fmt"
)

type PromAccount struct {
	client *Client
}

type OrderFunc func(Order)

func NewPromAccount(apiKey string) *PromAccount {
	acc := new(PromAccount)
	acc.client = NewClient(apiKey)
	return acc
}

func (acc *PromAccount) RequestOrders(f OrderFunc) (err error) {
	var orders Orders

	err = acc.client.Request("orders/list", map[string]string{ /*"status": "pending"*/ "limit": "2"}, &orders)
	fmt.Printf("%v", orders)
	fmt.Printf("%v", err)
	if err != nil {
		return
	}

	for _, order := range orders.Orders {
		f(order)
	}

	return
}
