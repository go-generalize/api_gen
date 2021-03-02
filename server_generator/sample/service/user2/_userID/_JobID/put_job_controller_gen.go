// Package _JobID ...
// generated version: devel
package _JobID

import (
	"github.com/go-generalize/api_gen/server_generator/sample/api_gen/props"
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
// @Param userID path string WIP:${isRequire} WIP:${description}
// @Param JobID path string WIP:${isRequire} WIP:${description}
// @Success 200 {object} PutJobResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /service/user2/{userID}/{JobID}/job [PUT]
func (p *PutJobController) PutJob(
	c echo.Context, req *PutJobRequest,
) (res *PutJobResponse, err error) {
	res = new(PutJobResponse)
	res.UserID = req.UserID
	res.JobID = req.JobID

	return res, nil
}
