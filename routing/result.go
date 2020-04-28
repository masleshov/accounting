package routing

import (
	"net/http"
)

// ExecutionResult represents result of executing some HTTP request
type ExecutionResult struct {
	status        int
	statusMessage string
	Response      []byte
}

// NewExecutionResult creates new instance of ExecutionResult type
func NewExecutionResult() *ExecutionResult {
	res := &ExecutionResult{}
	res.SetStatus(http.StatusOK)
	return res
}

// SetStatus sets status code and message to result of executioning HTTP request
func (res *ExecutionResult) SetStatus(statusCode int) {
	// checks that statusCode really exists in list of HTTP status codes
	if val := http.StatusText(statusCode); len(val) > 0 {
		res.status = statusCode
		res.statusMessage = val
	} else {
		res.status = http.StatusInternalServerError
		res.statusMessage = "Exception while setting status of result - unsupported status code"
	}
}

// GetStatus returns status code of execution result
func (res *ExecutionResult) GetStatus() int {
	return res.status
}

// GetStatusMessage returns status code of execution result
func (res *ExecutionResult) GetStatusMessage() string {
	return res.statusMessage
}
