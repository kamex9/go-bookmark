package utils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RetrievePathParamValue(key string, r *http.Request) (string, error) {
	vars := mux.Vars(r)
	value, ok := vars[key]
	if !ok {
		return "", fmt.Errorf("path parameter '%s' not defined", key)
	}
	return value, nil
}

func RetrievePathParamUintValue(key string, r *http.Request) (uint16, error) {
	value, err := RetrievePathParamValue(key, r)
	if err != nil {
		return 0, err
	}

	parsedValue, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("path parameter '%s' value '%d' is invalid: %v", key, parsedValue, err)
	}

	return uint16(parsedValue), nil
}
