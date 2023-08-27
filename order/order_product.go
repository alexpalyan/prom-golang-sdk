package order

type Product struct {
	ID            int                       `json:"id"`
	ExternalID    *string                   `json:"external_id"`
	Name          string                    `json:"name"`
	NameMultilang map[string]string         `json:"name_multilang"`
	Sku           string                    `json:"sku"`
	Price         string                    `json:"price"`
	Quantity      int                       `json:"quantity"`
	MeasureUnit   string                    `json:"measure_unit"`
	Image         string                    `json:"image"`
	URL           string                    `json:"url"`
	TotalPrice    string                    `json:"total_price"`
	CPACommission OrderProductCPACommission `json:"cpa_commission"`
}
