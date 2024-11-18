package utils

import (
	"encoding/json"
	"fmt"
	"go-bookmark/app/web/dto"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var l = GetLogger()

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

func WriteResponse(w http.ResponseWriter, data any, err *dto.Error) {
	var response dto.Response
	if err != nil {
		l.Error("Error occurred in processing", "error", *err)
		w.WriteHeader(http.StatusBadRequest)
		// Dataフィールドにはany型初期値のnilが設定される
		response = dto.Response{Error: err}
	} else {
		// Errorフィールドにはポインタ型初期値のnilが設定される
		// ※ポインタ型にしないと、ユーザー定義型の初期値である空オブジェクトが設定されてしまうためjson:omitemptyが効かなくなる
		response = dto.Response{Data: data}
	}

	json.NewEncoder(w).Encode(response)
}
