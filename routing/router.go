package routing

import (
	"accounting/accounting/service"
	"accounting/accounting/util"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// Router is class which contains routing map by requests and service methods
type Router struct {
	Tree RouteTree
}

// Instance is single instance of Router type
var Instance *Router

// NewRouter creates new instance of router
func NewRouter() *Router {
	if Instance == nil {
		Instance = &Router{
			Tree: NewRouteTree(),
		}
	}

	return Instance
}

// Route subscribe on route handlers
func (router *Router) Route() {
	for uri := range Instance.Tree {
		http.Handle(uri, router)
		log.Println("handled: " + uri)
	}
}

func (router *Router) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	for uri, route := range Instance.Tree {
		if !strings.Contains(request.RequestURI, uri) {
			continue
		}

		result := NewExecutionResult()
		if req := getRequest(route, request); req == nil {
			result.SetStatus(http.StatusNotImplemented)
			result.Response, _ = json.Marshal("Method with indicated type not implemented or not founded in route tree")
		} else {
			if err := validateRequestType(req, *request); err != nil {
				result.SetStatus(http.StatusBadRequest)
				result.Response, _ = json.Marshal(err.Error())
			} else {
				if res := execRequest(route, req); res.Error != nil {
					result.SetStatus(http.StatusInternalServerError)
					result.Response, _ = json.Marshal("Service exception: " + res.Error.Error())
				} else {
					result.Response = res.Object
				}
			}
		}

		response.Header().Set("Content-Type", "application/json")

		status := result.GetStatus()
		response.WriteHeader(status)

		if status != http.StatusOK {
			statusMessage, _ := json.Marshal(strconv.Itoa(status) + " " + result.GetStatusMessage())
			result.Response = append(statusMessage[:], result.Response[:]...)
		}

		response.Write(result.Response)
	}
}

func getRequest(route interface{}, request *http.Request) service.APIRequest {
	if _, ok := route.(service.GetMethod); ok {
		return service.NewGetRequest(*request)
	} else if _, ok := route.(service.PostMethod); ok {
		return service.NewPostRequest(*request)
	} else if _, ok := route.(service.PutMethod); ok {
		return service.NewPutRequest(*request)
	} else if _, ok := route.(service.DeleteMethod); ok {
		return service.NewDeleteRequest(*request)
	}

	return nil
}

func validateRequestType(request service.APIRequest, httpRequest http.Request) error {
	httpMethod := httpRequest.Method
	requestType := strings.ToUpper(strings.Replace(reflect.TypeOf(request).Name(), "Request", "", -1))

	if requestType != httpMethod {
		return errors.New("Wrong type of request, expected " + requestType)
	}

	return nil
}

func execRequest(route interface{}, request service.APIRequest) util.JSONObject {
	var result util.JSONObject

	if req, ok := request.(service.GetRequest); ok {
		method, _ := route.(service.GetMethod)
		result = method(req)
	} else if req, ok := request.(service.PostRequest); ok {
		method, _ := route.(service.PostMethod)
		result = method(req)
	} else if req, ok := request.(service.PutRequest); ok {
		method, _ := route.(service.PutMethod)
		result = method(req)
	} else if req, ok := request.(service.DeleteRequest); ok {
		method, _ := route.(service.DeleteMethod)
		result = method(req)
	} else {
		result = util.NewJSONObject(nil, errors.New("Service method with same type that request not founded"))
	}

	return result
}
