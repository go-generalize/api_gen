// Package static ...
package static

// GetStaticPageRequest ...
type GetStaticPageRequest struct{}

// GetStaticPageResponse ...
type GetStaticPageResponse struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
