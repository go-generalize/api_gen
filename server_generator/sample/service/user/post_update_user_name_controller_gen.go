// Package user ...
// generated version: devel
package user

import (
	"github.com/go-generalize/api_gen/server_generator/sample/apigen/props"
	"github.com/labstack/echo/v4"
)

// PostUpdateUserNameController ...
type PostUpdateUserNameController struct {
	*props.ControllerProps
}

// NewPostUpdateUserNameController ...
func NewPostUpdateUserNameController(cp *props.ControllerProps) *PostUpdateUserNameController {
	p := &PostUpdateUserNameController{
		ControllerProps: cp,
	}
	return p
}

// PostUpdateUserName ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param Name body string WIP:${isRequire} WIP:${description}
// @Success 200 {object} PostUpdateUserNameResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /service/user/update_user_name [POST]
func (p *PostUpdateUserNameController) PostUpdateUserName(
	c echo.Context, req *PostUpdateUserNameRequest,
) (res *PostUpdateUserNameResponse, err error) {
	// API Error Usage: github.com/go-generalize/api_gen/server_generator/sample/wrapper
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
