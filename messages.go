// messages
package prom

import (
	"fmt"
	"strconv"
	"time"
)

const (
	MessageStatusUnread  = "unread"
	MessageStatusRead    = "read"
	MessageStatusDeleted = "deleted"
)

type MessagesRequest struct {
	Status   string
	DateFrom time.Time
	DateTo   time.Time
	Limit    int
	LastId   int
}

type Message struct {
	Id             int    `json:"id"`
	DateCreated    string `json:"date_created"`
	ClientFullName string `json:"client_full_name"`
	Phone          string `json:"phone"`
	Message        string `json:"message"`
	Subject        string `json:"subject"`
	Status         string `json:"status"`
	ProductId      int    `json:"product_id"`
}

type MessagesResponse struct {
	Messages []Message `json:"messages"`
	Error    string    `json:"error"`
}

type MessageResponse struct {
	Message Message `json:"message"`
	Error   string  `json:"error"`
}

type SetMessageStatus struct {
	Status string `json:"status"`
	Ids    []int  `json:"ids"`
}

type SetMessageStatusResponse struct {
	ProcessedIds []int  `json:"processed_ids"`
	Error        string `json:"error"`
}

type MessageReply struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

type MessageReplyResponse struct {
	ProcessedIds []int  `json:"processed_ids"`
	Error        string `json:"error"`
}

func (c *Client) GetMessages(request MessagesRequest) (messages []Message, err error) {
	var (
		result MessagesResponse
		params = make(map[string]string)
	)

	if !request.DateFrom.IsZero() {
		params["date_from"] = request.DateFrom.Format(RequestDateFormat)
	}

	if !request.DateTo.IsZero() {
		params["date_to"] = request.DateTo.Format(RequestDateFormat)
	}

	if request.LastId > 0 {
		params["last_id"] = strconv.Itoa(request.LastId)
	}

	if request.Limit > 0 {
		params["limit"] = strconv.Itoa(request.LastId)
	}

	err = c.Get("/messages/list", params, &result)
	messages = result.Messages
	return
}

func (c *Client) GetMessage(id int) (message Message, err error) {
	var result MessageResponse

	err = c.Get("/messages/"+strconv.Itoa(id), nil, &result)
	if err != nil {
		err = fmt.Errorf("Error when request message: %s", err)
		return
	}

	if len(result.Error) > 0 {
		err = fmt.Errorf("Error when request message: %s", result.Error)
		return
	}
	message = result.Message
	return
}

func (c *Client) UpdateMessageStatus(status string, ids []int) (result SetMessageStatusResponse, err error) {

	request := SetMessageStatus{
		Status: status,
		Ids:    ids,
	}
	err = c.Post("/messages/set_status", request, &result)
	return
}

func (c *Client) ReplyMessage(id int, message string) (result MessageReplyResponse, err error) {
	request := MessageReply{
		Id:      id,
		Message: message,
	}

	err = c.Post("/messages/reply", request, &result)
	return
}
