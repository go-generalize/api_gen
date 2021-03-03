// Package user2 ...
// generated version: devel
package user2

import (
	"fmt"

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
// @Router /service/user2/update_user_name [POST]
func (p *PostUpdateUserNameController) PostUpdateUserName(
	_ echo.Context, _ *PostUpdateUserNameRequest,
) (res *PostUpdateUserNameResponse, err error) {
	return nil, fmt.Errorf("internal server error")
}
