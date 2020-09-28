// generated version: 0.4.0

package room

import (
	props "github.com/go-generalize/api_gen/server_generator/sample2/props"
	"github.com/labstack/echo/v4"
)

// GetRoomController ...
type GetRoomController struct {
	*props.ControllerProps
}

// NewGetRoomController ...
func NewGetRoomController(props *props.ControllerProps) *GetRoomController {
	g := &GetRoomController{
		ControllerProps: props,
	}
	return g
}

// GetRoom ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param eventID path string WIP:${isRequire} WIP:${description}
// @Success 200 {object} GetRoomResponse
// @Failure 400 {object} WIP
// @Router /api/event/{eventID}/room/room [GET]
func (g *GetRoomController) GetRoom(
	c echo.Context, req *GetRoomRequest,
) (res *GetRoomResponse, err error) {
	panic("require implements.")
}
