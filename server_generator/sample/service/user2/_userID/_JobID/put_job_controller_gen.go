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

// PutJob
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param userID path string WIP:${isRequire} WIP:${description}
// @Param JobID path string WIP:${isRequire} WIP:${description}
// @Success 200 {object} PutJobResponse
// @Failure 400 {object} WIP
// @Router /service/user2/{userID}/{JobID}/job [PUT]
func (p *PutJobController) PutJob(
	ctx context.Context, c echo.Context, req *PutJobRequest,
) (res *PutJobResponse, err error) {
	res = new(PutJobResponse)
	res.UserID = req.UserID
	res.JobID = req.JobID

	return res, nil
}
