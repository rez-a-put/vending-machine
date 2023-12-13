package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
	"vending-machine/model"
)

// setup response data
func ReturnResponse(w http.ResponseWriter, statusCode int, respMsg string, retData interface{}) {
	respData := &model.Response{
		Status:  strconv.Itoa(statusCode),
		Message: respMsg,
		Data:    retData,
	}

	// convert data into json and send as response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(respData)
	if err != nil {
		panic(err.Error())
	}
}
