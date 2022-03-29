package e2eutil

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func waitUntilServerReady(t *testing.T, addr string) {
	t.Helper()

	c := http.Client{
		Timeout: 1 * time.Second,
	}

	for i := 0; i < 10; i++ {
		_, err := c.Get(addr)

		if err == nil {
			return
		}

		time.Sleep(time.Second)
	}

	t.Fatal("server is not ready")
}

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

	addr = fmt.Sprintf("http://127.0.0.1:%d", port)

	waitUntilServerReady(t, addr)

	return addr, func() {
		close(ch)
	}
}
