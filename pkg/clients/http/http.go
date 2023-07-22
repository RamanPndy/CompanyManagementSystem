package http

import (
	"net/http"
)

// IRequest ...
type IRequest interface {
	Get(r *InnerRequest) (*http.Response, error)
	Post(r *InnerRequest) (*http.Response, error)
}

// NewRequest Creates an instance if a request
func NewRequest() IRequest {
	return &Request{}
}

// Request , is an invoker struct for the interface
type Request struct {
}

// InnerRequest contains a method to perform an HTTP request
type InnerRequest struct {
	Req *http.Request
}

// Get makes a http get request
func (r *Request) Get(req *InnerRequest) (*http.Response, error) {
	req.Req.Method = "GET"
	return do(req)
}

// Post makes a http post request
func (r *Request) Post(req *InnerRequest) (*http.Response, error) {
	req.Req.Method = "POST"
	return do(req)
}

func do(r *InnerRequest) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(r.Req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
