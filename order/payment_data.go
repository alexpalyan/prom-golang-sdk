package order

type PaymentData struct {
	Type           string        `json:"type"`
	Status         PaymentStatus `json:"status"`
	StatusModified interface{}   `json:"status_modified"`
}
