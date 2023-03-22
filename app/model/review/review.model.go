package reviewmodel

import (
	"time"
)

const (
	StatusInactive = 0
	StatusActive   = 1
)

type Review struct {
	ID          int64     `json:"id"             db:"id"`
	UserID      int64     `json:"user_id"        db:"user_id"`
	BookID      int64     `json:"book_id"        db:"book_id"`
	Rate        int       `json:"rate"           db:"rate"`
	Content     string    `json:"content"        db:"content"`
	Status      int       `json:"status"         db:"status"`
	CreatedTime time.Time `json:"created_time"   db:"created_time"`
}
