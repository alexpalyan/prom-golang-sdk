package prom

// import (
// 	"fmt"
// )

type PromAccount struct {
	client *Client
}

func NewPromAccount(apiKey string) *PromAccount {
	acc := new(PromAccount)
	acc.client = NewClient(apiKey)
	return acc
}
