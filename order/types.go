package order

type Status string

func (s Status) String() string {
	return string(s)
}

type Source string

func (s Source) String() string {
	return string(s)
}

type CancellationReason string

func (c CancellationReason) String() string {
	return string(c)
}

type DeliveryProvider string

func (d DeliveryProvider) String() string {
	return string(d)
}

type DeliveryProviderType string

func (d DeliveryProviderType) String() string {
	return string(d)
}

type PaymentStatus string

func (p PaymentStatus) String() string {
	return string(p)
}
