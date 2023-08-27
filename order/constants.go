package order

const (
	StatusPending   Status = "pending"
	StatusReceived  Status = "received"
	StatusDelivered Status = "delivered"
	StatusCanceled  Status = "canceled"
	StatusDraft     Status = "draft"
	StatusPaid      Status = "paid"

	SourcePortal         Source = "portal"
	SourceCompanySite    Source = "company_site"
	SourceCompanyCabinet Source = "company_cabinet"
	SourceMobileApp      Source = "mobile_app"
	SourceBigl           Source = "bigl"

	CancellationReasonNotAvailable         CancellationReason = "not_available"
	CancellationReasonPriceChanged         CancellationReason = "price_changed"
	CancellationReasonBuyersRequest        CancellationReason = "buyers_request"
	CancellationReasonNotEnoughFields      CancellationReason = "not_enough_fields"
	CancellationReasonDuplicate            CancellationReason = "duplicate"
	CancellationReasonInvalidPhoneNumber   CancellationReason = "invalid_phone_number"
	CancellationReasonLessThanMinimalPrice CancellationReason = "less_than_minimal_price"
	CancellationReasonAnother              CancellationReason = "another"

	DeliveryProviderNovaPoshta   DeliveryProvider = "nova_poshta"
	DeliveryProviderJustin       DeliveryProvider = "justin"
	DeliveryProviderDeliveryAuto DeliveryProvider = "delivery_auto"
	DeliveryProviderUkrposhta    DeliveryProvider = "ukrposhta"

	DeliveryProviderDataTypeW2W DeliveryProviderType = "W2W"
	DeliveryProviderDataTypeW2D DeliveryProviderType = "W2D"
	DeliveryProviderDataTypeD2W DeliveryProviderType = "D2W"
	DeliveryProviderDataTypeD2D DeliveryProviderType = "D2D"

	PaymentStatusPaid     PaymentStatus = "paid"
	PaymentStatusUnPaid   PaymentStatus = "unpaid"
	PaymentStatusRefunded PaymentStatus = "refunded"
	PaymentStatusPaidOut  PaymentStatus = "paid_out"
)
