package service

// GetArticleRequest ...
type GetArticleRequest struct {
	ID int
}

// GetArticleResponse ...
type GetArticleResponse struct {
	ID    int
	Group []string
	Body  string
}
