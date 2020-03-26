// client
package prom

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	apiKey string
	apiUrl string
}

func NewClient(apiKey string) *Client {
	c := &Client{
		apiUrl: defaultApiUrl,
		apiKey: apiKey,
	}
	return c
}

const defaultApiUrl = "https://my.prom.ua/api/v1"

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

	if resp.StatusCode >= 400 {
		err = fmt.Errorf("Error when request: %s", resp.Status)
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err = json.Unmarshal(respBody, v); err != nil {
		return
	}

	return
}

func (c *Client) Get(route string, params map[string]string, v interface{}) (err error) {
	req, err := http.NewRequest(http.MethodGet, c.apiUrl+route, nil)
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

	req, err := http.NewRequest(http.MethodPost, c.apiUrl+route, r)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	err = c.Request(req, v)
	return
}
