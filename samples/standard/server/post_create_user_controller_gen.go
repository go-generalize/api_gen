// Package sample ...
// generated version: devel
package sample

import (
	"github.com/go-generalize/api_gen/samples/standard/server/props"
	"github.com/labstack/echo/v4"
)

// PostCreateUserController ...
type PostCreateUserController struct {
	*props.ControllerProps
}

// NewPostCreateUserController ...
func NewPostCreateUserController(cp *props.ControllerProps) *PostCreateUserController {
	p := &PostCreateUserController{
		ControllerProps: cp,
	}
	return p
}

// PostCreateUser ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param ID body string false ""
// @Param Password body string false ""
// @Param Gender body integer false ""
// @Param Birthday body time.Time false ""
// @Param Roles body []*Role false ""
// @Success 200 {object} PostCreateUserResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /create_user [POST]
func (p *PostCreateUserController) PostCreateUser(
	c echo.Context, req *PostCreateUserRequest,
) (res *PostCreateUserResponse, err error) {
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
func (p *PostCreateUserController) AutoBind() bool {
	return true
}
