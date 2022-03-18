package e2eutil

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// StartServer starts a server with given ctrlInitializer.
func StartServer[Props any, Ctrl any](
	t *testing.T,
	p *Props,
	ctrlInitializer func(p *Props, e *echo.Echo) *Ctrl,
) (addr string, stop func()) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	ctrlInitializer(p, e)

	ch := make(chan struct{})

	port := rand.Intn(50000) + 1001

	go func() {
		e.Start(":" + strconv.Itoa(port))
		<-ch
		e.Shutdown(context.Background())
	}()

	return fmt.Sprintf("http://127.0.0.1:%d", port), func() {
		close(ch)
	}
}
