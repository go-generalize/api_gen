// generated version: unknown
package _userID

import (
	"github.com/go-generalize/api_gen/server_generator/sample/props"
	"github.com/labstack/echo/v4"
)

// GetUserJobGetController ...
type GetUserJobGetController struct {
	*props.ControllerProps
}

// NewGetUserJobGetController ...
func NewGetUserJobGetController(props *props.ControllerProps) *GetUserJobGetController {
	g := &GetUserJobGetController{
		ControllerProps: props,
	}
	return g
}

// GetUserJobGet ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param userID path string WIP:${isRequire} WIP:${description}
// @Success 200 {object} GetUserJobGetResponse
// @Failure 400 {object} WIP
// @Router /service/user2/{userID}/user_job_get [GET]
func (g *GetUserJobGetController) GetUserJobGet(
	c echo.Context, req *GetUserJobGetRequest,
) (res *GetUserJobGetResponse, err error) {
	panic("require implements.")
}
