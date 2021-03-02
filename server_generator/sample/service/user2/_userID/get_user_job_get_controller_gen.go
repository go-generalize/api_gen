// Package _userID ...
// generated version: devel
package _userID

import (
	"github.com/go-generalize/api_gen/server_generator/sample/api_gen/props"
	"github.com/labstack/echo/v4"
)

// GetUserJobGetController ...
type GetUserJobGetController struct {
	*props.ControllerProps
}

// NewGetUserJobGetController ...
func NewGetUserJobGetController(cp *props.ControllerProps) *GetUserJobGetController {
	g := &GetUserJobGetController{
		ControllerProps: cp,
	}
	return g
}

// GetUserJobGet ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param userID path string WIP:${isRequire} WIP:${description}
// @Success 200 {object} GetUserJobGetResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /service/user2/{userID}/user_job_get [GET]
func (g *GetUserJobGetController) GetUserJobGet(
	c echo.Context, req *GetUserJobGetRequest,
) (res *GetUserJobGetResponse, err error) {
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
