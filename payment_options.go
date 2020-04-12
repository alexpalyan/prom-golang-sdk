// payment_options
package prom

type PaymentOption struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PaymentOptionsResponse struct {
	PaymentOptions []PaymentOption `json:"payment_options"`
	Error          string          `json:"string"`
}

func (c *Client) GetPaymentOptions() (paymentOptions []PaymentOption, err error) {
	var result PaymentOptionsResponse
	err = c.Get("/payment_options/list", nil, &result)
	paymentOptions = result.PaymentOptions
	return
}
