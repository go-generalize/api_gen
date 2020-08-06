package service

import (
	"context"

	"github.com/labstack/echo/v4"
)

type GetArticleController struct {
}

func NewGetArticleController() *GetArticleController {
	g := &GetArticleController{}
	return g
}

// GetArticle
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param ID query integer WIP:${isRequire} WIP:${description}
// @Success 200 {object} GetArticleResponse
// @Failure 400 {object} WIP
// @Router /service/article [GET]
func (g *GetArticleController) GetArticle(
	ctx context.Context, c echo.Context, req *GetArticleRequest,
) (res *GetArticleResponse, err error) {
	return &GetArticleResponse{
		ID:    req.ID * 2,
		Group: []string{"a", "b", "c"},
		Body:  "hogehoge",
	}, nil
}
