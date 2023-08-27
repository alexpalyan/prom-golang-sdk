package order

import (
	"time"
)

type Order struct {
	ID                   int                   `json:"id"`
	DateCreated          time.Time             `json:"date_created"`
	DateModified         time.Time             `json:"date_modified"`
	ClientFirstName      string                `json:"client_first_name"`
	ClientSecondName     *string               `json:"client_second_name"`
	ClientLastName       string                `json:"client_last_name"`
	ClientID             int                   `json:"client_id"`
	Email                *string               `json:"email"`
	Phone                string                `json:"phone"`
	DeliveryOption       *DeliveryOption       `json:"delivery_option"`
	DeliveryAddress      string                `json:"delivery_address"`
	DeliveryProviderData *DeliveryProviderData `json:"delivery_provider_data"`
	DeliveryCost         float64               `json:"delivery_cost"`
	PaymentOption        *PaymentOption        `json:"payment_option"`
	PaymentData          *PaymentData          `json:"payment_data"`
	Price                string                `json:"price"`
	FullPrice            string                `json:"full_price"`
	ClientNotes          string                `json:"client_notes"`
	Products             []Product             `json:"products"`
	Status               Status                `json:"status"`
	Source               string                `json:"source"`
}
