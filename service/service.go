package service

import (
	"accounting/accounting/data/repository"
	"accounting/accounting/util"
	"io/ioutil"
	"log"
	"net/http"
)

// APIRequest is some request to API
type APIRequest interface{}

// Request represents a some HTTP request
type Request struct {
	APIRequest
	Body []byte
}

// GetRequest represents HTTP GET request
type GetRequest struct {
	Request
}

// PostRequest represents HTTP POST request
type PostRequest struct {
	Request
}

// PutRequest represents HTTP PUT request
type PutRequest struct {
	Request
}

// DeleteRequest represents HTTP DELETE request
type DeleteRequest struct {
	Request
}

// GetMethod represents GET API method
type GetMethod = func(GetRequest) util.JSONObject

// PostMethod represents POST API method
type PostMethod = func(PostRequest) util.JSONObject

// PutMethod represents PUT API method
type PutMethod = func(PutRequest) util.JSONObject

// DeleteMethod represents DELETE API method
type DeleteMethod = func(DeleteRequest) util.JSONObject

// NewGetRequest creates a new GET API request instance by HTTP request
func NewGetRequest(req http.Request) GetRequest {
	return GetRequest{
		Request: *newRequest(req),
	}
}

// NewPostRequest creates a new POST API request instance by HTTP request
func NewPostRequest(req http.Request) PostRequest {
	return PostRequest{
		Request: *newRequest(req),
	}
}

// NewPutRequest creates a new PUT API request instance by HTTP request
func NewPutRequest(req http.Request) PutRequest {
	return PutRequest{
		Request: *newRequest(req),
	}
}

// NewDeleteRequest creates a new DELETE API request instance by HTTP request
func NewDeleteRequest(req http.Request) DeleteRequest {
	return DeleteRequest{
		Request: *newRequest(req),
	}
}

// ValidateBodyIsEmpty checks that GET request doesn't have no ony parameter
func (request *GetRequest) ValidateBodyIsEmpty() {
	if len(request.Body) > 0 {
		log.Fatalf("Method shouldn't receive no one parameters")
	}
}

func newRequest(req http.Request) *Request {
	return &Request{
		Body: readHTTPBody(req),
	}
}

func readHTTPBody(req http.Request) []byte {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Error while reading request body: " + err.Error())
	}

	return body
}

func toJSON(data repository.DBResult) util.JSONObject {
	return util.NewJSONObject(data.Data, data.Error)
}

func errorToJSON(err error) util.JSONObject {
	return util.NewJSONObject(nil, err)
}
