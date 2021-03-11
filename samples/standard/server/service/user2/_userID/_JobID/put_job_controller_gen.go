// Package _JobID ...
// generated version: devel
package _JobID

import (
	"github.com/go-generalize/api_gen/samples/standard/server/props"
	"github.com/labstack/echo/v4"
)

// PutJobController ...
type PutJobController struct {
	*props.ControllerProps
}

// NewPutJobController ...
func NewPutJobController(cp *props.ControllerProps) *PutJobController {
	p := &PutJobController{
		ControllerProps: cp,
	}
	return p
}

// PutJob ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param userID path string true ""
// @Param JobID path string true ""
// @Success 200 {object} PutJobResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /service/user2/{userID}/{JobID}/job [PUT]
func (p *PutJobController) PutJob(
	c echo.Context, req *PutJobRequest,
) (res *PutJobResponse, err error) {
	// API Error Usage: github.com/go-generalize/api_gen/samples/standard/server/wrapper
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

// AutoBind - use echo.Bind
func (p *PutJobController) AutoBind() bool {
	return true
}
