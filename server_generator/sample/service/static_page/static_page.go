package static

type GetStaticPageRequest struct{}

type GetStaticPageResponse struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
