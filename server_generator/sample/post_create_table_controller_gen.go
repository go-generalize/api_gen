// Package sample ...
// generated version: devel
package sample

import (
	"github.com/go-generalize/api_gen/server_generator/sample/props"
	"github.com/go-generalize/api_gen/server_generator/sample/service/table"
	"github.com/labstack/echo/v4"
)

// PostCreateTableController ...
type PostCreateTableController struct {
	*props.ControllerProps
}

// NewPostCreateTableController ...
func NewPostCreateTableController(props *props.ControllerProps) *PostCreateTableController {
	p := &PostCreateTableController{
		ControllerProps: props,
	}
	return p
}

// PostCreateTable ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param ID body string WIP:${isRequire} WIP:${description}
// @Param Text body string WIP:${isRequire} WIP:${description}
// @Success 200 {object} PostCreateTableResponse
// @Failure 400 {object} WIP
// @Router /create_table [POST]
func (p *PostCreateTableController) PostCreateTable(
	c echo.Context, req *PostCreateTableRequest,
) (res *PostCreateTableResponse, err error) {
	id := p.TestKey

	res = &PostCreateTableResponse{
		ID:      id,
		Payload: table.Table{},
	}

	return res, nil
}
