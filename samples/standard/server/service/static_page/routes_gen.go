// Code generated by server_generator. DO NOT EDIT.
// generated version: devel

package static

import (
	"io"
	"log"
	"net/http"

	"github.com/go-generalize/api_gen/samples/standard/server/props"
	"github.com/go-generalize/api_gen/samples/standard/server/wrapper"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
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
	router.GET("static_page", r.GetStaticPage(p))
	return r
}

// GetStaticPage ...
func (r *Routes) GetStaticPage(p *props.ControllerProps) echo.HandlerFunc {
	i := NewGetStaticPageController(p)

	b, ok := (interface{})(i).(interface{ AutoBind() bool })
	bindable := !ok || b.AutoBind()

	return func(c echo.Context) error {
		var (
			req  *GetStaticPageRequest
			werr *wrapper.APIError
		)

		if bindable {
			req = new(GetStaticPageRequest)
			if err := c.Bind(req); err != nil {
				log.Printf("failed to JSON binding(/service/static_page/static_page): %+v", err)
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return c.JSON(werr.Status, werr.Body)
				}
				return err
			}
		}
		res, err := i.GetStaticPage(c, req)
		if err != nil {
			if xerrors.As(err, &werr) {
				log.Printf("%+v", werr)
				return c.JSON(werr.Status, werr.Body)
			}
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

// IGetStaticPageController ...
type IGetStaticPageController interface {
	GetStaticPage(c echo.Context, req *GetStaticPageRequest) (res *GetStaticPageResponse, err error)
}
