// order
package prom

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/alexpalyan/prom-golang-sdk/order"
)

type ErrorResponse struct {
	Text string
}

func (e *ErrorResponse) UnmarshalJSON(b []byte) (err error) {
	var text string
	if err = json.Unmarshal(b, &text); err != nil {
		return
	}

	e.Text = text

	return
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("Error response received during request: %s", e.Text)
}

type OrdersResponse struct {
	Orders []order.Order  `json:"orders"`
	Error  *ErrorResponse `json:"error"`
}

type OrderResponse struct {
	Order order.Order    `json:"order"`
	Error *ErrorResponse `json:"error"`
}

type OrdersRequest struct {
	Status   *order.Status
	DateFrom time.Time
	DateTo   time.Time
	Limit    *int
	LastId   *int
}

type SetOrderStatus struct {
	Status             order.Status `json:"status"`
	IDs                []int        `json:"ids"`
	CancellationReason string       `json:"cancellation_reason,omitempty"`
	CancellationText   string       `json:"cancellation_text,omitempty"`
}

type OrdersSetStatusResponse struct {
	ProcessedIDs []int          `json:"processed_ids"`
	Error        *ErrorResponse `json:"error"`
}
