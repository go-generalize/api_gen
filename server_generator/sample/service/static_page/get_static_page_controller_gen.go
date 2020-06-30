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

// GetStaticPage
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Success 200 {object} GetStaticPageResponse
// @Failure 400 {object} WIP
// @Router /service/static_page/static_page [GET]
func (g *GetStaticPageController) GetStaticPage(
	ctx context.Context, c echo.Context, req *GetStaticPageRequest,
) (res *GetStaticPageResponse, err error) {
	panic("require implements.")
}
