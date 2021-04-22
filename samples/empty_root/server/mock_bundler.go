// +build mock

// Code generated by api_gen. DO NOT EDIT.
// generated version: devel
package controller

import (
	"net/http"

	types_bar_8619c483 "github.com/go-generalize/api_gen/samples/empty_root/api/foo/bar"
	ctrl_bar_8bd5d3ca "github.com/go-generalize/api_gen/samples/empty_root/server/mock/controller/foo/bar"
	apierror "github.com/go-generalize/api_gen/samples/empty_root/server/pkg/apierror"
	props "github.com/go-generalize/api_gen/samples/empty_root/server/props"
	echo "github.com/labstack/echo/v4"
	xerrors "golang.org/x/xerrors"
)

func addRoutes(e *echo.Echo, p *props.ControllerProps) {
	add := func(method, path string, handler func(c echo.Context) (interface{}, error)) {
		e.Add(method, path, func(c echo.Context) error {
			var werr *apierror.APIError

			res, err := handler(c)

			if err != nil {
				if xerrors.As(err, &werr) {
					c.Logger().Errorf("%+v", werr)
					return c.JSON(werr.Status, werr.Body)
				}
				return err
			}
			if res == nil {
				return nil
			}

			return c.JSON(http.StatusOK, res)
		})
	}

	{
		ctrl := ctrl_bar_8bd5d3ca.NewPostUserController(p)

		add("POST", "/foo/bar/user", func(c echo.Context) (interface{}, error) {
			var werr *apierror.APIError

			req := new(types_bar_8619c483.PostUserRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/foo/bar/user): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, err
			}

			res, err := ctrl.PostUser(c, req)

			if err != nil {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, xerrors.Errorf("PostUser returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

}
