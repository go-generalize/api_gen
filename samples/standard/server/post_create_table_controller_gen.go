// Package sample ...
// generated version: devel
package sample

import (
	"github.com/go-generalize/api_gen/samples/standard/server/props"
	"github.com/labstack/echo/v4"
)

// PostCreateTableController ...
type PostCreateTableController struct {
	*props.ControllerProps
}

// NewPostCreateTableController ...
func NewPostCreateTableController(cp *props.ControllerProps) *PostCreateTableController {
	p := &PostCreateTableController{
		ControllerProps: cp,
	}
	return p
}

// PostCreateTable ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param ID body string false ""
// @Param Text body string false ""
// @Param Flag body Flag false ""
// @Param map body map[Flag]Flag false ""
// @Success 200 {object} PostCreateTableResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /create_table [POST]
func (p *PostCreateTableController) PostCreateTable(
	c echo.Context, req *PostCreateTableRequest,
) (res *PostCreateTableResponse, err error) {
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
func (p *PostCreateTableController) AutoBind() bool {
	return true
}
