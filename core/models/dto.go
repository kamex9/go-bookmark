package models

import (
	"encoding/json"
	"time"
)

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
