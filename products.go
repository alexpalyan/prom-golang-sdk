// product
package prom

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	ProductSellingTypeRetail    = "retail"
	ProductSellingTypeWholesale = "wholesale"
	ProductSellingTypeUniversal = "universal"
	ProductSellingTypeService   = "service"

	ProductPresenceAvailable    = "available"
	ProductPresenceNotAvailable = "not_available"
	ProductPresenceOrder        = "order"
	ProductPresenceService      = "service"
	ProductPresenceWaiting      = "waiting"

	ProductStatusOnDisplay          = "on_display"
	ProductStatusDraft              = "draft"
	ProductStatusDeleted            = "deleted"
	ProductStatusNotOnDisplay       = "not_on_display"
	ProductStatusEditingRequired    = "editing_required"
	ProductStatusApprovalPending    = "approval_pending"
	ProductStatusDeletedByModerator = "deleted_by_moderator"

	ProductDiscountTypeAmount  = "amount"
	ProductDiscountTypePercent = "percent"

	ProductMarkMissingProductAsNone         = "none"
	ProductMarkMissingProductAsNotAvailable = "not_available"
	ProductMarkMissingProductAsNotOnDisplay = "not_on_display"
	ProductMarkMissingProductAsDeleted      = "deleted"

	ProductUpdatedFieldsName            = "name"
	ProductUpdatedFieldsSku             = "sku"
	ProductUpdatedFieldsPrice           = "price"
	ProductUpdatedFieldsImagesUrls      = "images_urls"
	ProductUpdatedFieldsPresence        = "presence"
	ProductUpdatedFieldsQuantityInStock = "quantity_in_stock"
	ProductUpdatedFieldsDescription     = "description"
	ProductUpdatedFieldsGroup           = "group"
	ProductUpdatedFieldsKeywords        = "keywords"
	ProductUpdatedFieldsAttributes      = "attributes"
	ProductUpdatedFieldsDiscount        = "discount"
	ProductUpdatedFieldsLabels          = "labels"
	ProductUpdatedFieldsGtin            = "gtin"
	ProductUpdatedFieldsMpn             = "mpn"

	ProductImportRequestStatusSuccess = "success"
	ProductImportRequestStatusError   = "error"

	ProductImportResultStatusSuccess = "SUCCESS"
	ProductImportResultStatusPartial = "PARTIAL"
	ProductImportResultStatusFatal   = "FATAL"
)

type ProductPrice struct {
	Price                float64 `json:"price"`
	MinimumOrderQuantity float64 `json:"minimum_order_quantity"`
}

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
	Prices    []ProductPrice `json:"prices,omitempty"`
	MainImage string         `json:"main_image"`
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

type ProductsRequest struct {
	Limit   int
	LastId  int
	GroupId int
}

type ProductsResponse struct {
	Products []Product `json:"products"`
	Error    string    `json:"error"`
}

type ProductResponse struct {
	Product Product `json:"product"`
	Error   string  `json:"error"`
}

type ProductEdit struct {
	Id           int            `json:"id"`
	Presence     string         `json:"presence"`
	PresenceSure bool           `json:"presence_sure"`
	Price        float64        `json:"price"`
	Status       string         `json:"status"`
	Prices       []ProductPrice `json:"prices,omitempty"`
	Discount     *Discount      `json:"discount"`
	Name         string         `json:"name"`
	Keywords     string         `json:"keywords"`
	Description  string         `json:"description"`
}

type ProductEditByExternalId struct {
	Id              string         `json:"id"`
	Presence        string         `json:"presence"`
	PresenceSure    bool           `json:"presence_sure"`
	Price           float64        `json:"price"`
	Status          string         `json:"status"`
	Prices          []ProductPrice `json:"prices,omitempty"`
	Discount        *Discount      `json:"discount"`
	Name            string         `json:"name"`
	Keywords        string         `json:"keywords"`
	Description     string         `json:"description"`
	QuantityInStock int            `json:"quantity_in_stock"`
}

type ProductEditResponse struct {
	ProcessedIds []int                  `json:"processed_ids"`
	Errors       map[string]interface{} `json:"errors"`
	Error        string                 `json:"error"`
}

type ImportProductURL struct {
	Url                  string   `json:"url"`
	ForceUpdate          bool     `json:"force_update"`
	OnlyAvailable        bool     `json:"only_available"`
	MarkMissingProductAs string   `json:"mark_missing_product_as"`
	UpdatedFields        []string `json:"updated_fields"`
}

type ImportProductWithFile struct {
	File string `json:"file"`
	Data struct {
		ForceUpdate          bool     `json:"force_update"`
		OnlyAvailable        bool     `json:"only_available"`
		MarkMissingProductAs string   `json:"mark_missing_product_as"`
		UpdatedFields        []string `json:"updated_fields"`
	}
}

type ProductImportResponse struct {
	Id      string `json:"id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ImportProductStatus struct {
	Status          string `json:"status"`
	NotChanged      int    `json:"not_changed"`
	Updated         int    `json:"updated"`
	NotInFile       int    `json:"not_in_fle"`
	Imported        int    `json:"imported"`
	Created         int    `json:"created"`
	Actualized      int    `json:"actualized"`
	CreatedActive   int    `json:"created_active"`
	CreatedHidden   int    `json:"created_hidden"`
	Total           int    `json:"total"`
	WithErrorsCount int    `json:"with_errors_count"`
	Errors          []struct {
		Download       interface{} `json:"download"`
		Store_file     interface{} `json:"store_file"`
		Validation     interface{} `json:"validation"`
		Import         interface{} `json:"import"`
		DownloadImages interface{} `json:"download_images"`
	} `json:"errors"`
	Message string `json:"message"`
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

func (c *Client) GetProducts(request ProductsRequest) (products []Product, err error) {
	var (
		result ProductsResponse
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
		result = ProductsResponse{}
		if limit > 0 && limit <= MaxLimit {
			params["limit"] = strconv.Itoa(limit)
		} else if limit > MaxLimit {
			params["limit"] = strconv.Itoa(MaxLimit)
		}

		err = c.Get("/products/list", params, &result)
		if err != nil {
			return nil, fmt.Errorf("Error when request products: %s", err)
		}

		if len(result.Products) > 0 {
			products = append(products, result.Products...)
			params["last_id"] = strconv.Itoa(result.Products[len(result.Products)-1].Id)
		}
		if limit <= MaxLimit || len(products) < MaxLimit {
			break
		}
		limit = limit - MaxLimit
	}

	return
}

func (c *Client) GetProduct(id int) (product Product, err error) {
	var result ProductResponse

	err = c.Get("/product/"+strconv.Itoa(id), nil, result)
	if err != nil {
		err = fmt.Errorf("Error when request product: %s", err)
		return
	}

	if len(result.Error) > 0 {
		err = fmt.Errorf("Error when request product: %s", result.Error)
		return
	}
	product = result.Product
	return
}

func (c *Client) GetProductByExternalId(id string) (product Product, err error) {
	var result ProductResponse

	err = c.Get("/product/by_external_id/"+url.QueryEscape(id), nil, result)
	if err != nil {
		err = fmt.Errorf("Error when request product: %s", err)
		return
	}

	if len(result.Error) > 0 {
		err = fmt.Errorf("Error when request product: %s", result.Error)
		return
	}
	product = result.Product
	return
}

func (c *Client) UpdateProducts(products []ProductEdit) (updatedIds []int, err error) {
	var result ProductEditResponse
	err = c.Post("/products/edit", products, &result)
	if err != nil {
		err = fmt.Errorf("Error when update product: %s", err)
		return
	}
	if len(result.Error) > 0 {
		err = fmt.Errorf("Error when update product: %s", result.Error)
		return
	}
	if len(result.Errors) > 0 {
		err = fmt.Errorf("Error when update product: %#v", result.Errors)
		return
	}
	updatedIds = result.ProcessedIds
	return
}

func (c *Client) UpdateProductsByExternalId(products []ProductEditByExternalId) (updatedIds []int, err error) {
	var result ProductEditResponse
	err = c.Post("/products/edit_by_external_id", products, &result)
	if err != nil {
		err = fmt.Errorf("Error when update product: %s", err)
		return
	}
	if len(result.Error) > 0 {
		err = fmt.Errorf("Error when update product: %s", result.Error)
		return
	}
	if len(result.Errors) > 0 {
		err = fmt.Errorf("Error when update product: %#v", result.Errors)
		return
	}
	updatedIds = result.ProcessedIds
	return
}

func (c *Client) ImportProductsByUrl(request ImportProductURL) (response ProductImportResponse, err error) {
	err = c.Post("/products/import_url", request, &response)
	return
}

func (c *Client) ImportProductsByFile(request ImportProductURL) (response ProductImportResponse, err error) {
	err = c.Post("/products/import_file", request, &response)
	return
}

func (c *Client) GetProductsImportStatus(id int) (response ImportProductStatus, err error) {
	err = c.Get("/products/import/status/"+strconv.Itoa(id), nil, &response)
	return
}
