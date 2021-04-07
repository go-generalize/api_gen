// Package static ...
// generated version: devel
package static

import (
	"github.com/go-generalize/api_gen/samples/standard/server/props"
	"github.com/labstack/echo/v4"
)

// GetStaticPageController ...
type GetStaticPageController struct {
	*props.ControllerProps
}

// NewGetStaticPageController ...
func NewGetStaticPageController(cp *props.ControllerProps) *GetStaticPageController {
	g := &GetStaticPageController{
		ControllerProps: cp,
	}
	return g
}

// GetStaticPage ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Success 200 {object} GetStaticPageResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /service/static_page/static_page [GET]
func (g *GetStaticPageController) GetStaticPage(
	c echo.Context, req *GetStaticPageRequest,
) (res *GetStaticPageResponse, err error) {
	// API Error Usage: github.com/go-generalize/api_gen/samples/standard/server/wrapper
	//
	// return nil, wrapper.NewAPIError(http.StatusBadRequest)
	//
	// return nil, wrapper.NewAPIError(http.StatusBadRequest).SetError(err)
	//
	// body := map[string]interface{}{
	// 	"code": http.StatusBadRequest,
	// 	"message": "invalid request parameter.",
	// }
	// return nil, wrapper.NewAPIError(http.StatusBadRequest, body).SetError(err)
	panic("require implements.") // FIXME require implements.
}

// AutoBind - use echo.Bind
func (g *GetStaticPageController) AutoBind() bool {
	return true
}
