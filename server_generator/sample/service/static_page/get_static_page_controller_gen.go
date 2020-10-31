// Package static ...
// generated version: unknown
package static

import (
	"github.com/go-generalize/api_gen/server_generator/sample/props"
	"github.com/labstack/echo/v4"
)

// GetStaticPageController ...
type GetStaticPageController struct {
	*props.ControllerProps
}

// NewGetStaticPageController ...
func NewGetStaticPageController(props *props.ControllerProps) *GetStaticPageController {
	g := &GetStaticPageController{
		ControllerProps: props,
	}
	return g
}

// GetStaticPage ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Success 200 {object} GetStaticPageResponse
// @Failure 400 {object} WIP
// @Router /service/static_page/static_page [GET]
func (g *GetStaticPageController) GetStaticPage(
	c echo.Context, req *GetStaticPageRequest,
) (res *GetStaticPageResponse, err error) {
	panic("require implements.")
}
