// Package room ...
// generated version: devel
package room

import (
	"github.com/go-generalize/api_gen/server_generator/sample2/props"
	"github.com/labstack/echo/v4"
)

// GetRoomController ...
type GetRoomController struct {
	*props.ControllerProps
}

// NewGetRoomController ...
func NewGetRoomController(cp *props.ControllerProps) *GetRoomController {
	g := &GetRoomController{
		ControllerProps: cp,
	}
	return g
}

// GetRoom ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param roomID path string WIP:${isRequire} WIP:${description}
// @Param eventID path string WIP:${isRequire} WIP:${description}
// @Success 200 {object} GetRoomResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /api/event/{eventID}/room/{roomID} [GET]
func (g *GetRoomController) GetRoom(
	c echo.Context, req *GetRoomRequest,
) (res *GetRoomResponse, err error) {
	// API Error Usage: github.com/go-generalize/api_gen/server_generator/sample2/wrapper
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
