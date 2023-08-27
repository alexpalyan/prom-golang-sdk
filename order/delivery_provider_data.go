package order

type DeliveryProviderData struct {
	Provider             DeliveryProvider     `json:"nova_poshta"`
	Type                 DeliveryProviderType `json:"type"`
	SenderWarehouseId    string               `json:"sender_warehouse_id"`
	RecipientWarehouseId string               `json:"recipient_warehouse_id"`
	DeclarationNumber    string               `json:"declaration_number"`
}
