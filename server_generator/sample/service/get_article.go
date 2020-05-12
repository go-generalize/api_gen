package service

type GetArticleRequest struct {
	ID int
}

type GetArticleResponse struct {
	ID    int
	Group []string
	Body  string
}
