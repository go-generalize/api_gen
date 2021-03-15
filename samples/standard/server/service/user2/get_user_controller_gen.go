// Package user2 ...
// generated version: devel
package user2

import (
	"github.com/go-generalize/api_gen/samples/standard/server/props"
	"github.com/labstack/echo/v4"
)

// GetUserController ...
type GetUserController struct {
	*props.ControllerProps
}

// NewGetUserController ...
func NewGetUserController(cp *props.ControllerProps) *GetUserController {
	g := &GetUserController{
		ControllerProps: cp,
	}
	return g
}

// GetUser ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param userID path string true ""
// @Param search_request query string false ""
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /service/user2/{userID} [GET]
func (g *GetUserController) GetUser(
	c echo.Context, req *GetUserRequest,
) (res *GetUserResponse, err error) {
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
func (g *GetUserController) AutoBind() bool {
	return true
}
