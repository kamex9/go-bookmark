package models

import (
	"encoding/json"
	"time"
)

// 共通のレスポンス構造体
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewError(code string, err error) *Error {
	return &Error{
		Code:    code,
		Message: err.Error(),
	}
}

type Response struct {
	Data  any    `json:"data"`
	Error *Error `json:"error,omitempty"`
}

type CustomTime time.Time

// MarshalJSON implements the json.Marshaler interface.
// It formats the time according to the specified layout.
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(ct).Format(time.DateTime))
}

type BookmarkPostRequest struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

type BookmarkResponse struct {
	ID        uint16     `json:"id"`
	URL       string     `json:"url"`
	Title     string     `json:"title"`
	CreatedAt CustomTime `json:"created_at"`
}
