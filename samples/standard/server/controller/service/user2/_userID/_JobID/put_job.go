// Package controller is for
// generated version: devel
package controller

import (
	types "github.com/go-generalize/api_gen/samples/standard/api/service/user2/_userID/_JobID"
	props "github.com/go-generalize/api_gen/samples/standard/server/props"
	"github.com/labstack/echo/v4"
)

// PutJobController ...
type PutJobController interface {
	PutJob(c echo.Context, req *types.PutJobRequest) (res *types.PutJobResponse, err error)
}

type putJobController struct {
	*props.ControllerProps
}

// NewPutJobController ...
func NewPutJobController(cp *props.ControllerProps) PutJobController {
	return &putJobController{
		ControllerProps: cp,
	}
}

// PutJob - PUT job
func (ctrl *putJobController) PutJob(
	c echo.Context, req *types.PutJobRequest,
) (res *types.PutJobResponse, err error) {
	// return nil, apierror.NewAPIError(http.StatusBadRequest)
	//
	// return nil, apierror.NewAPIError(http.StatusBadRequest).SetError(err)
	//
	// body := map[string]interface{}{
	// 	"code": http.StatusBadRequest,
	// 	"message": "invalid request parameter.",
	// }
	// return nil, apierror.NewAPIError(http.StatusBadRequest, body).SetError(err)
	panic("require implements.") // FIXME require implements.
}

// AutoBind - use echo.Bind
func (ctrl *putJobController) AutoBind() bool {
	return true
}
