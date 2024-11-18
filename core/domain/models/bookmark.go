package models

import (
	"time"
)

type Bookmark struct {
	ID        string
	URL       string
	Title     string
	CreatedAt time.Time
}
