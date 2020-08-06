package sample

import (
	"context"

	"github.com/go-generalize/api_gen/server_generator/sample/service/table"
	"github.com/labstack/echo/v4"
)

type PostCreateTableController struct {
}

func NewPostCreateTableController() *PostCreateTableController {
	p := &PostCreateTableController{}
	return p
}

// PostCreateTable
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param ID body string WIP:${isRequire} WIP:${description}
// @Param Text body string WIP:${isRequire} WIP:${description}
// @Success 200 {object} PostCreateTableResponse
// @Failure 400 {object} WIP
// @Router /create_table [POST]
func (p *PostCreateTableController) PostCreateTable(
	ctx context.Context, c echo.Context, req *PostCreateTableRequest,
) (res *PostCreateTableResponse, err error) {
	id := ctx.Value(testKey).(string)

	res = &PostCreateTableResponse{
		ID:      id,
		Payload: table.Table{},
	}

	return res, nil
}
