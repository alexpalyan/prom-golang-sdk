// client
package prom

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/alexpalyan/prom-golang-sdk/order"
	paymentoption "github.com/alexpalyan/prom-golang-sdk/payment_option"
)

type Client struct {
	apiKey string
	apiURL string
}

func NewClient(apiKey string) *Client {
	c := &Client{
		apiURL: defaultAPIURL,
		apiKey: apiKey,
	}
	return c
}

const defaultAPIURL = "https://my.prom.ua/api/v1"

const (
	RequestDateTimeFormat = "2006-01-02T15:04:05"
	RequestDateFormat     = "02.01.2006"
	DiscountDateFormat    = "02.01.2006"
	MaxLimit              = 100
)

func (c *Client) Request(req *http.Request, v interface{}) (err error) {
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		err = &ClientError{
			Code: resp.StatusCode,
			Body: string(respBody),
		}
		return
	}

	if err = json.Unmarshal(respBody, v); err != nil {
		return
	}

	return
}

func (c *Client) Get(route string, params map[string]string, v interface{}) (err error) {
	req, err := http.NewRequest(http.MethodGet, c.apiURL+route, nil)
	if err != nil {
		return
	}
	q := req.URL.Query()
	for k, vl := range params {
		q.Set(k, vl)
	}
	req.URL.RawQuery = q.Encode()
	err = c.Request(req, v)
	return
}

func (c *Client) Post(route string, data interface{}, v interface{}) (err error) {
	r, w := io.Pipe()

	go func() {
		defer w.Close()
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			return
		}
	}()

	req, err := http.NewRequest(http.MethodPost, c.apiURL+route, r)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	err = c.Request(req, v)
	return
}

func (c *Client) GetPaymentOptions() (paymentOptions []paymentoption.PaymentOption, err error) {
	var result PaymentOptionsResponse
	err = c.Get("/payment_options/list", nil, &result)
	paymentOptions = result.PaymentOptions
	return
}

func (c *Client) GetOrders(request OrdersRequest) (orders []order.Order, err error) {
	var (
		result OrdersResponse
		params = make(map[string]string)
	)

	if request.Status != nil {
		params["status"] = request.Status.String()
	}

	if !request.DateFrom.IsZero() {
		params["date_from"] = request.DateFrom.Format(RequestDateTimeFormat)
	}

	if !request.DateTo.IsZero() {
		params["date_to"] = request.DateTo.Format(RequestDateTimeFormat)
	}

	if request.LastId != nil {
		params["last_id"] = strconv.Itoa(*request.LastId)
	}

	if request.Limit != nil && *request.Limit > MaxLimit {
		return nil, fmt.Errorf("Limit must be less than %d", MaxLimit)
	} else if request.Limit != nil && *request.Limit == 0 {
		return nil, fmt.Errorf("Limit must be greater than 0")
	} else if request.Limit != nil {
		params["limit"] = strconv.Itoa(*request.Limit)
	}

	result = OrdersResponse{}

	err = c.Get("/orders/list", params, &result)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	if len(result.Orders) > 0 {
		orders = append(orders, result.Orders...)
		params["last_id"] = strconv.Itoa(result.Orders[len(result.Orders)-1].ID)
	}

	return
}

func (c *Client) GetOrder(id int) (*order.Order, error) {
	var result OrderResponse

	err := c.Get("/orders/"+strconv.Itoa(id), nil, &result)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &result.Order, nil
}

func (c *Client) UpdateOrdersStatus(s SetOrderStatus) ([]int, error) {
	var result OrdersSetStatusResponse

	err := c.Post("/orders/set_status", s, &result)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return result.ProcessedIDs, nil
}
