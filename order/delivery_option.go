package order

type DeliveryOption struct {
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	ShippingService interface{} `json:"shipping_service"`
}
