// Package foo ...
// generated version: devel
package foo

import (
	"github.com/go-generalize/api_gen/samples/empty_root/server/props"
	"github.com/labstack/echo/v4"
)

// PostUserController ...
type PostUserController struct {
	*props.ControllerProps
}

// NewPostUserController ...
func NewPostUserController(cp *props.ControllerProps) *PostUserController {
	p := &PostUserController{
		ControllerProps: cp,
	}
	return p
}

// PostUser ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Success 200 {object} PostUserResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /foo/bar/user [POST]
func (p *PostUserController) PostUser(
	c echo.Context, req *PostUserRequest,
) (res *PostUserResponse, err error) {
	// API Error Usage: github.com/go-generalize/api_gen/samples/empty_root/server/wrapper
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
func (p *PostUserController) AutoBind() bool {
	return true
}
