// payment_options
package prom

import (
	paymentoption "github.com/alexpalyan/prom-golang-sdk/payment_option"
)

type PaymentOptionsResponse struct {
	PaymentOptions []paymentoption.PaymentOption `json:"payment_options"`
	Error          string                        `json:"string"`
}
