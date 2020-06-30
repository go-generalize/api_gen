package user2

import (
	"context"

	"github.com/labstack/echo/v4"
)

type PostUpdateUserNameController struct {
}

func NewPostUpdateUserNameController() *PostUpdateUserNameController {
	p := &PostUpdateUserNameController{}
	return p
}

// PostUpdateUserName
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param
// @Success 200 {object} PostUpdateUserNameResponse
// @Failure 400 {object} WIP
// @Router /service/user2/update_user_name [POST]
func (p *PostUpdateUserNameController) PostUpdateUserName(
	ctx context.Context, c echo.Context, req *PostUpdateUserNameRequest,
) (res *PostUpdateUserNameResponse, err error) {
	panic("require implements.")
}
