package prom

import "fmt"

type ClientError struct {
	Code int
	Body string
}

func (e *ClientError) Error() string {
	return fmt.Sprintf("Error when request PROM (Code: %d, Body: %s)", e.Code, e.Body)
}

type RequestError struct {
	reqType   string
	originErr error
}

func NewRequestError(t string, err error) *RequestError {
	return &RequestError{
		reqType:   t,
		originErr: err,
	}
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("Error request %s: %s", e.reqType, e.originErr)
}

func (e *RequestError) GetOriginError() error {
	return e.originErr
}

type ResponseError struct {
	reqType   string
	originErr string
}

func NewResponseError(t string, e string) *ResponseError {
	return &ResponseError{
		reqType:   t,
		originErr: e,
	}
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("Error when request %s: %s", e.reqType, e.originErr)
}
