// products_test
package prom

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetProducts(t *testing.T) {
	require := require.New(t)

	ts := httptest.NewServer(
		http.HandlerFunc(
			CreateServerDummy(http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/products/list":

					response, _ := os.ReadFile("testdata/products_list.json")
					w.Write(response)
				default:
					w.WriteHeader(http.StatusNotFound)
					return
				}
			}),
		),
	)
	defer ts.Close()

	c := &Client{
		apiURL: ts.URL,
		apiKey: "authorizedTestKey",
	}

	request := ProductsRequest{}
	products, err := c.GetProducts(request)

	require.NoError(err)
	require.NotNil(products)

	require.Equal(1, len(products))
}
