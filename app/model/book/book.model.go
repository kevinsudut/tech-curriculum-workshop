package bookmodel

import (
	"time"
)

type Book struct {
	ID          int64     `json:"id"             db:"id"`
	Title       string    `json:"title"          db:"title"`
	Subtitle    string    `json:"subtitle"       db:"subtitle"`
	Description string    `json:"description"    db:"description"`
	ISBN        string    `json:"isbn"           db:"isbn"`
	Author      string    `json:"author"         db:"author"`
	Published   time.Time `json:"published"      db:"published"`
	Publisher   string    `json:"publisher"      db:"publisher"`
	Pages       int       `json:"pages"          db:"pages"`
}
