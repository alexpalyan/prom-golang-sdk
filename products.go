// product
package prom

import (
	"fmt"
	"math"
	"strconv"
)

type Product struct {
	Id                   int       `json:"id"`
	ExternalId           string    `json:"external_id"`
	Name                 string    `json:"name"`
	Sku                  string    `json:"sku"`
	Keywords             string    `json:"keywords"`
	Description          string    `json:"description"`
	SellingType          string    `json:"selling_type"`
	Presence             string    `json:"presence"`
	PresenceSure         bool      `json:"presence_sure"`
	Price                float64   `json:"price"`
	MinimumOrderQuantity float64   `json:"minimum_order_quantity"`
	Discount             *Discount `json:"discount"`
	Currency             string    `json:"currency"`
	Group                struct {
		Id   int    `json:"id"`
		Name string `json:"string"`
	} `json:"group"`
	Category struct {
		Id      int    `json:"id"`
		Caption string `json:"caption"`
	} `json:"category"`
	Prices []struct {
		Price                float64 `json:"price"`
		MinimumOrderQuantity float64 `json:"minimum_order_quantity"`
	} `json:"prices"`
	MainImage string `json:"main_image"`
	Images    []struct {
		Url          string `json:"url"`
		ThumbnailUrl string `json:"thumbnail_url"`
		Id           int    `json:"id"`
	} `json:"images"`
	Status string `json:"status"`
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

type ProductEdit struct {
	Id           int     `json:"id"`
	Presence     string  `json:"presence"`
	PresenceSure bool    `json:"presence_sure"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
	Prices       []struct {
		Price                float64 `json:"price"`
		MinimumOrderQuantity float64 `json:"minimum_order_quantity"`
	} `json:"prices,omitempty"`
	Discount    *Discount `json:"discount"`
	Name        string    `json:"name"`
	Keywords    string    `json:"keywords"`
	Description string    `json:"description"`
}

type ProductEditResponce struct {
	ProcessedIds []int       `json:"processed_ids"`
	Errors       interface{} `json:"errors"`
	Error        interface{} `json:"error"`
}

const (
	DiscountTypeAmount  = "amount"
	DiscountTypePercent = "percent"
)

func (acc *PromAccount) GetProducts(request ProductsRequest) (products []Product, err error) {
	var (
		result Products
		params map[string]string = make(map[string]string)
	)

	if request.GroupId >= 0 {
		params["group_id"] = strconv.Itoa(request.GroupId)
	}

	if request.LastId > 0 {
		params["last_id"] = strconv.Itoa(request.LastId)
	}
	limit := request.Limit

	for {
		if limit > 0 && limit <= 100 {
			params["limit"] = strconv.Itoa(request.Limit)
		} else if limit > 100 {
			params["limit"] = "100"
			limit = limit - 100
		}

		err = acc.client.Get("products/list", params, &result)
		if err != nil {
			return nil, fmt.Errorf("Error when request products: %s", result.Error)
		}

		if len(result.Products) > 0 {
			products = append(products, result.Products...)
			params["last_id"] = strconv.Itoa(result.Products[len(result.Products)-1].Id)
		}
		if len(result.Products) <= int(math.Min(100.0, float64(limit))) {
			break
		}
	}

	return
}

func (acc *PromAccount) UpdateProducts(product []ProductEdit) (err error) {

	var result ProductEditResponce
	acc.client.Post("products/edit", product, &result)
	fmt.Printf("%#v", result)
	return nil
}

func NewProductEdit(product Product) (result ProductEdit) {
	result = ProductEdit{
		Id:           product.Id,
		Presence:     product.Presence,
		PresenceSure: product.PresenceSure,
		Price:        product.Price,
		Status:       product.Status,
		Prices:       append(product.Prices[:0:0], product.Prices...),
		Name:         product.Name,
		Keywords:     product.Keywords,
		Description:  product.Description,
	}

	if product.Discount != nil {
		result.Discount = &Discount{
			Value:     product.Discount.Value,
			Type:      product.Discount.Type,
			DateStart: product.Discount.DateStart,
			DateEnd:   product.Discount.DateEnd,
		}
	}

	return
}
