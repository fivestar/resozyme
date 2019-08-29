package model

import "time"

// Article is a model represents an article.
type Article struct {
	ID      int64
	Title   string
	PubDate time.Time
}
