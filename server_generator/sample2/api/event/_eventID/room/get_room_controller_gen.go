// generated version: 0.3.4
package room

import (
	"context"

	"github.com/labstack/echo/v4"
)

type GetRoomController struct {
}

func NewGetRoomController() *GetRoomController {
	g := &GetRoomController{}
	return g
}

// GetRoom
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param eventID path string WIP:${isRequire} WIP:${description}
// @Success 200 {object} GetRoomResponse
// @Failure 400 {object} WIP
// @Router /api/event/{eventID}/room/room [GET]
func (g *GetRoomController) GetRoom(
	ctx context.Context, c echo.Context, req *GetRoomRequest,
) (res *GetRoomResponse, err error) {
	panic("require implements.")
}
