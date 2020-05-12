package sample

import (
	"github.com/labstack/echo/v4"
)

type PostCreateTableController struct {
}

func NewPostCreateTableController() *PostCreateTableController {
	p := &PostCreateTableController{}
	return p
}

func (p *PostCreateTableController) PostCreateTable(c echo.Context,
	req *PostCreateTableRequest) (
	res *PostCreateTableResponse,
	err error) {
	panic("require implements.")
}
