package service

import (
	"context"

	"github.com/labstack/echo/v4"
)

type GetInfoController struct {
}

func NewGetInfoController() *GetInfoController {
	g := &GetInfoController{}
	return g
}

func (g *GetInfoController) GetInfo(ctx context.Context,
	c echo.Context,
	req *GetInfoRequest) (
	res *GetInfoResponse,
	err error) {
	panic("require implements.")
}
