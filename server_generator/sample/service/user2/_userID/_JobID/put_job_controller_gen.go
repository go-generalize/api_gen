package _JobID

import (
	"context"

	"github.com/labstack/echo/v4"
)

type PutJobController struct {
}

func NewPutJobController() *PutJobController {
	p := &PutJobController{}
	return p
}

func (p *PutJobController) PutJob(
	ctx context.Context, c echo.Context, req *PutJobRequest,
) (res *PutJobResponse, err error) {
	res = new(PutJobResponse)
	res.UserID = req.UserID
	res.JobID = req.JobID

	return res, nil
}
