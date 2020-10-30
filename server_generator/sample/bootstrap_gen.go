// Code generated by server_generator. DO NOT EDIT.
// generated version: unknown

package sample

import (
	"io"
	"log"
	"net/http"

	"github.com/go-generalize/api_gen/server_generator/sample/props"
	"github.com/go-generalize/api_gen/server_generator/sample/service"
	serviceStaticPage "github.com/go-generalize/api_gen/server_generator/sample/service/static_page"
	serviceUser "github.com/go-generalize/api_gen/server_generator/sample/service/user"
	serviceUser2 "github.com/go-generalize/api_gen/server_generator/sample/service/user2"
	serviceUser2UserID "github.com/go-generalize/api_gen/server_generator/sample/service/user2/_userID"
	serviceUser2UserIDJobID "github.com/go-generalize/api_gen/server_generator/sample/service/user2/_userID/_JobID"
	"github.com/labstack/echo/v4"
)

// MiddlewareList ...
type MiddlewareList []*MiddlewareSet

// MiddlewareMap ...
type MiddlewareMap map[string][]echo.MiddlewareFunc

// MiddlewareSet ...
type MiddlewareSet struct {
	Path           string
	MiddlewareFunc []echo.MiddlewareFunc
}

// ToMap ...
func (m MiddlewareList) ToMap() MiddlewareMap {
	mf := make(map[string][]echo.MiddlewareFunc)
	for _, middleware := range m {
		mf[middleware.Path] = middleware.MiddlewareFunc
	}
	return mf
}

// Bootstrap ...
func Bootstrap(p *props.ControllerProps, e *echo.Echo, middlewareList MiddlewareList, opts ...io.Writer) {
	if len(opts) > 0 {
		if w := opts[0]; w != nil {
			log.SetOutput(w)
		}
	}

	middleware := middlewareList.ToMap()

	// error handling
	e.Use(func(before echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			defer func() {
				recoverErr := recover()
				if recoverErr == nil {
					return
				}
				if httpErr, ok := recoverErr.(*echo.HTTPError); ok {
					err = c.JSON(httpErr.Code, httpErr.Message)
				}
				log.Printf("panic: %#v", recoverErr)
				err = c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": "internal server error.",
				})
			}()

			err = before(c)
			return err
		}
	})

	rootGroup := e.Group("")
	setMiddleware(rootGroup, "/", middleware)
	NewRoutes(p, rootGroup, opts...)
	serviceGroup := rootGroup.Group("service/")
	setMiddleware(serviceGroup, "/service/", middleware)
	service.NewRoutes(p, serviceGroup, opts...)
	serviceStaticPageGroup := serviceGroup.Group("static_page/")
	setMiddleware(serviceStaticPageGroup, "/service/static_page/", middleware)
	serviceStaticPage.NewRoutes(p, serviceStaticPageGroup, opts...)
	serviceUserGroup := serviceGroup.Group("user/")
	setMiddleware(serviceUserGroup, "/service/user/", middleware)
	serviceUser.NewRoutes(p, serviceUserGroup, opts...)
	serviceUser2Group := serviceGroup.Group("user2/")
	setMiddleware(serviceUser2Group, "/service/user2/", middleware)
	serviceUser2.NewRoutes(p, serviceUser2Group, opts...)
	serviceUser2UserIDGroup := serviceUser2Group.Group(":userID/")
	setMiddleware(serviceUser2UserIDGroup, "/service/user2/:userID/", middleware)
	serviceUser2UserID.NewRoutes(p, serviceUser2UserIDGroup, opts...)
	serviceUser2UserIDJobIDGroup := serviceUser2UserIDGroup.Group(":JobID/")
	setMiddleware(serviceUser2UserIDJobIDGroup, "/service/user2/:userID/:JobID/", middleware)
	serviceUser2UserIDJobID.NewRoutes(p, serviceUser2UserIDJobIDGroup, opts...)
}

func setMiddleware(group *echo.Group, path string, list MiddlewareMap) {
	if ms, ok := list[path]; ok {
		for _, m := range ms {
			group.Use(m)
		}
	}
}
