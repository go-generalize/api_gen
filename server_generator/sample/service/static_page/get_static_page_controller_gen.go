package static

import (
	"context"

	"github.com/labstack/echo/v4"
)

type GetStaticPageController struct {
}

func NewGetStaticPageController() *GetStaticPageController {
	g := &GetStaticPageController{}
	return g
}

func (g *GetStaticPageController) GetStaticPage(
	_ context.Context, _ echo.Context, _ *GetStaticPageRequest,
) (res *GetStaticPageResponse, err error) {
	panic("require implements.")
}
