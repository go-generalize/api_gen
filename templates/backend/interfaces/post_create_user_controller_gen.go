package interfaces

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostCreateUserController struct {
}

func NewPostCreateUserController() *PostCreateUserController {
	p := &PostCreateUserController{}
	return p
}

func (p *PostCreateUserController) PostCreateUser(
	ctx context.Context, c echo.Context, req *PostCreateUserRequest,
) (res *PostCreateUserResponse, err error) {
	// reqCtx := c.Request().Context()

	if len(req.Password) == 0 {
		return nil, c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": "password is not set",
			},
		)
	}

	return &PostCreateUserResponse{
		Status:      true,
		Message:     "ok",
		CreatedType: CreatedTypeMember,
	}, nil
}
