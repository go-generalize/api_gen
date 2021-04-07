// Code generated by server_generator. DO NOT EDIT.
// generated version: devel

package foo

import (
	"io"
	"log"
	"net/http"

	"github.com/go-generalize/api_gen/samples/empty_root/server/props"
	"github.com/go-generalize/api_gen/samples/empty_root/server/wrapper"
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
	router.POST("user", r.PostUser(p))
	return r
}

// PostUser ...
func (r *Routes) PostUser(p *props.ControllerProps) echo.HandlerFunc {
	i := NewPostUserController(p)

	b, ok := (interface{})(i).(interface{ AutoBind() bool })
	bindable := !ok || b.AutoBind()

	return func(c echo.Context) error {
		var (
			req  *PostUserRequest
			werr *wrapper.APIError
		)

		if bindable {
			req = new(PostUserRequest)
			if err := c.Bind(req); err != nil {
				log.Printf("failed to JSON binding(/foo/bar/user): %+v", err)
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
		res, err := i.PostUser(c, req)
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

// IPostUserController ...
type IPostUserController interface {
	PostUser(c echo.Context, req *PostUserRequest) (res *PostUserResponse, err error)
}
