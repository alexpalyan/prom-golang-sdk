package prom

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/alexpalyan/prom-golang-sdk/order"
	"github.com/alexpalyan/prom-golang-sdk/utils"
	"github.com/stretchr/testify/require"
)

func TestGetOrders(t *testing.T) {
	require := require.New(t)
	ts := httptest.NewServer(http.HandlerFunc(
		CreateServerDummy(http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/orders/list":

				if r.URL.Query().Get("status") == "invalid" {
					response := []byte(`{"error": "Incorrect status value"}`)
					w.Write(response)
					return
				}

				response, _ := os.ReadFile("testdata/orders_list.json")
				w.Write(response)

			default:
				w.WriteHeader(http.StatusNotFound)
				return
			}
		})))
	defer ts.Close()

	c := &Client{
		apiURL: ts.URL,
		apiKey: "authorizedTestKey",
	}
	request := OrdersRequest{}
	orders, err := c.GetOrders(request)
	require.NoError(err)
	require.NotNil(orders)

	require.Equal(2, len(orders))
	ord := orders[0]
	require.Equal(123, ord.ID)
	require.Equal(order.StatusDelivered, ord.Status)

	request.Status = utils.Ptr(order.Status("invalid"))
	orders, err = c.GetOrders(request)
	require.Error(err)
	require.Nil(orders)
}
