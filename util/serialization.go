package util

import (
	"encoding/json"
)

// JSONSerializer serialize interface{} to JSON
type JSONSerializer interface {
	Serialize(value interface{}, err error)
}

// JSONObject is type which represents JSON response and error
type JSONObject struct {
	Object []byte
	Error  error
}

// NewJSONObject creates new instance of JSONObject type
func NewJSONObject(value interface{}, err error) JSONObject {
	res := &JSONObject{
		Object: serialize(value),
		Error:  err,
	}

	return *res
}

func serialize(value interface{}) []byte {
	res, _ := json.Marshal(value)
	return res
}
