// client
package prom

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	c := new(Client)
	c.apiKey = apiKey
	return c
}

const apiUrl = "https://my.prom.ua/api/v1/"

func (c *Client) Request(route string, params map[string]string, v interface{}) (err error) {

	req := &http.Request{
		Method: http.MethodGet,
		Header: http.Header{
			"Authorization": {"Bearer " + c.apiKey},
		},
	}

	req.URL, _ = url.Parse(apiUrl + route)
	q := req.URL.Query()
	for k, vl := range params {
		q.Set(k, vl)
	}
	req.URL.RawQuery = q.Encode()

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
