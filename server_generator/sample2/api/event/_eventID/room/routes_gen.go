// Code generated by server_generator. DO NOT EDIT.
// generated version: unknown

package room

import (
	"io"
	"log"
	"net/http"

	"github.com/go-generalize/api_gen/server_generator/sample2/props"
	"github.com/labstack/echo/v4"
)

// Routes ...
type Routes struct {
	router *echo.Group
}

// NewRoutes ...
func NewRoutes(p *props.ControllerProps, router *echo.Group, opts ...io.Writer) *Routes {
	if len(opts) > 0 {
		if w := opts[0]; w != nil {
			log.SetOutput(w)
		}
	}
	r := &Routes{
		router: router,
	}
	router.GET(":roomID", r.GetRoom(p))
	return r
}

// GetRoom ...
func (r *Routes) GetRoom(p *props.ControllerProps) echo.HandlerFunc {
	i := NewGetRoomController(p)
	return func(c echo.Context) error {
		req := new(GetRoomRequest)
		if err := c.Bind(req); err != nil {
			log.Printf("failed to JSON binding(/api/event/{eventID}/room/{roomID}): %+v", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.GetRoom(c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

// IGetRoomController ...
type IGetRoomController interface {
	GetRoom(c echo.Context, req *GetRoomRequest) (res *GetRoomResponse, err error)
}
