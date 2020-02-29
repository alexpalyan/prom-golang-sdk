// product
package prom

import (
	"strconv"
	// "encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Discount    *Discount `json:"discount"`
	Currency    string    `json:"currency"`
	Status      string    `json:"status"`
}

type Discount struct {
	Value     float64 `json:"value"`
	Type      string  `json:"type"`
	DateStart string  `json:"date_start"`
	DateEnd   string  `json:"date_end"`
}

type Products struct {
	Products []Product `json:"products"`
	Error    string    `json:"error"`
}

type ProductsRequest struct {
	Limit   int
	LastId  int
	GroupId int
}

func (acc *PromAccount) GetProducts(request ProductsRequest) (products []Product, err error) {
	var (
		result Products
		params map[string]string = make(map[string]string)
	)

	if request.GroupId >= 0 {
		params["group_id"] = strconv.Itoa(request.GroupId)
	}

	if request.Limit > 0 && request.Limit <= 100 {
		params["limit"] = strconv.Itoa(request.Limit)
	} else if request.Limit > 100 {
		params["limit"] = "100"
	}

	if request.LastId > 0 {
		params["last_id"] = strconv.Itoa(request.LastId)
	}

	for {
		err = acc.client.Request("products/list", params, &result)
		if err != nil {
			return nil, fmt.Errorf("Error when request products: %s", result.Error)
		}

		if len(result.Products) > 0 {
			products = append(products, result.Products...)
			params["last_id"] = strconv.Itoa(result.Products[len(result.Products)-1].Id)
		} else {
			break
		}
	}

	return
}
