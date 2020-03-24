// order
package prom

import (
	"fmt"
	"strconv"
	"time"
)

const (
	OrderStatusPending   = "pending"
	OrderStatusReceived  = "received"
	OrderStatusDelivered = "delivered"
	OrderStatusCanceled  = "canceled"
	OrderStatusDraft     = "draft"
	OrderStatusPaid      = "paid"

	OrderSourcePortal         = "portal"
	OrderSourceCompanySite    = "company_site"
	OrderSourceCompanyCabinet = "company_cabinet"
	OrderSourceMobileApp      = "mobile_app"
	OrderSourceBigl           = "bigl"

	OrderCancellationReasonNotAvailable         = "not_available"
	OrderCancellationReasonPriceChanged         = "price_changed"
	OrderCancellationReasonBuyersRequest        = "buyers_request"
	OrderCancellationReasonNotEnoughFields      = "not_enough_fields"
	OrderCancellationReasonDuplicate            = "duplicate"
	OrderCancellationReasonInvalidPhoneNumber   = "invalid_phone_number"
	OrderCancellationReasonLessThanMinimalPrice = "less_than_minimal_price"
	OrderCancellationReasonAnother              = "another"
)

type Order struct {
	Id               int    `json:"id"`
	DateCreated      string `json:"date_created"`
	ClientFirstName  string `json:"client_first_name"`
	ClientSecondName string `json:"client_second_name"`
	ClientLastName   string `json:"client_last_name"`
	ClientId         int    `json:"client_id"`
	ClientNotes      string `json:"client_notes"`
	Products         []struct {
		Id          int     `json:"id"`
		ExternalId  string  `json:"external_id"`
		Image       string  `json:"image"`
		Quantity    float32 `json:"quantity"`
		Price       string  `json:"price"`
		Url         string  `json:"url"`
		Name        string  `json:"name"`
		TotalPrice  string  `json:"total_price"`
		MeasureUnit string  `json:"measure_unit"`
		Sku         string  `json:"sku"`
	} `json:"products"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Price          string `json:"price"`
	DeliveryOption struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"delivery_option"`
	DeliveryAddress string `json:"delivery_address"`
	PaymentOption   struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	Status string `json:"status"`
	Source string `json:"source"`
}

type OrdersResponse struct {
	Orders []Order `json:"orders"`
	Error  string  `json:"error"`
}

type OrderResponse struct {
	Order Order  `json:"order"`
	Error string `json:"error"`
}

type OrdersRequest struct {
	Status   string
	DateFrom time.Time
	DateTo   time.Time
	Limit    int
	LastId   int
}

type SetOrderStatus struct {
	Status             string `json:"status"`
	Ids                []int  `json:"ids"`
	CancellationReason string `json:"cancellation_reason,omitempty"`
	CancellationText   string `json:"cancellation_text,omitempty"`
}

type OrdersSetStatusResponse struct {
	ProcessedIds []int `json:"processed_ids"`
	Error        string
}

func (c *Client) GetOrders(request OrdersRequest) (orders []Order, err error) {
	var (
		result OrdersResponse
		params map[string]string = make(map[string]string)
	)

	if len(request.Status) > 0 {
		params["status"] = request.Status
	}

	if !request.DateFrom.IsZero() {
		params["date_from"] = request.DateFrom.Format("2006-01-02T15:04:05")
	}

	if !request.DateTo.IsZero() {
		params["date_from"] = request.DateTo.Format("2006-01-02T15:04:05")
	}

	if request.LastId > 0 {
		params["last_id"] = strconv.Itoa(request.LastId)
	}
	limit := request.Limit

	for {
		result = OrdersResponse{}
		if limit > 0 && limit <= MaxLimit {
			params["limit"] = strconv.Itoa(limit)
		} else if limit > MaxLimit {
			params["limit"] = strconv.Itoa(MaxLimit)
		}

		err = c.Get("/orders/list", params, &result)
		if err != nil {
			return nil, fmt.Errorf("Error when request orders: %s", err)
		}
		if len(result.Error) > 0 {
			return nil, fmt.Errorf("Error when request orders: %s", result.Error)
		}

		if len(result.Orders) > 0 {
			orders = append(orders, result.Orders...)
			params["last_id"] = strconv.Itoa(result.Orders[len(result.Orders)-1].Id)
		}
		if limit <= MaxLimit || len(orders) < MaxLimit {
			break
		}
		limit = limit - MaxLimit
	}

	return
}

func (c *Client) GetOrder(id int) (order Order, err error) {
	var result OrderResponse

	err = c.Get("/orders/"+strconv.Itoa(id), nil, &result)
	if err != nil {
		err = fmt.Errorf("Error when request order: %s", err)
		return
	}

	if len(result.Error) > 0 {
		err = fmt.Errorf("Error when request order: %s", result.Error)
		return
	}

	order = result.Order
	return
}

func (c *Client) UpdateOrdersStatus(s SetOrderStatus) (ids []int, err error) {
	var result OrdersSetStatusResponse
	err = c.Post("/orders/set_status", s, &result)
	if err != nil {
		err = fmt.Errorf("Error when set_status orders: %s", err)
		return
	}
	if len(result.Error) > 0 {
		err = fmt.Errorf("Error when set_status orders: %s", result.Error)
		return
	}
	ids = result.ProcessedIds
	return
}
