// Package service ...
package service

import "time"

// GetArticleRequest ...
type GetArticleRequest struct {
	ID int
}

// GetArticleResponse ...
type GetArticleResponse struct {
	ID          int
	Group       []string
	Body        string
	RequestTime time.Time
}
