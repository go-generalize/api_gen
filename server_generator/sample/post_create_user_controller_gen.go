// Package sample ...
// generated version: devel
package sample

import (
	"fmt"

	"github.com/go-generalize/api_gen/server_generator/sample/props"
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
// @Param ID body string WIP:${isRequire} WIP:${description}
// @Param Password body string WIP:${isRequire} WIP:${description}
// @Param Gender body integer WIP:${isRequire} WIP:${description}
// @Param Birthday body time.Time WIP:${isRequire} WIP:${description}
// @Param Roles body []*Role WIP:${isRequire} WIP:${description}
// @Success 200 {object} PostCreateUserResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /create_user [POST]
func (p *PostCreateUserController) PostCreateUser(
	_ echo.Context, req *PostCreateUserRequest,
) (res *PostCreateUserResponse, err error) {
	return &PostCreateUserResponse{
		Status:      true,
		Message:     fmt.Sprintf("password is %s", req.Password),
		CreatedType: CreatedTypeMember,
	}, nil
}
