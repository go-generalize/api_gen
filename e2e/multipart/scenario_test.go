//go:build internal
// +build internal

package scenario_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/go-gneralize/api_gen/v2/e2e/e2eutil"
	"github.com/go-gneralize/api_gen/v2/e2e/multipart/controller"
)

func TestTypeScript(t *testing.T) {
	rootDir := os.Getenv("E2E_API_GEN_ROOT_DIR")

	e2eutil.RunCommandWithModifier(t, func(cmd *exec.Cmd) {
		cmd.Dir = "./tests/ts"
	}, filepath.Join(rootDir, "node_modules/.bin/webpack"), "--config", "webpack.config.js")

	su := e2eutil.NewSeleniumUtil(t)
	defer su.Close()

	addr, stop := e2eutil.StartServer(t, nil, controller.NewControllers)
	defer stop()

	su.MustExecuteScript(
		t,
		e2eutil.ReadFile(t, "tests/ts/dist/main.js"),
		[]interface{}{addr},
	)
}

func TestDart(t *testing.T) {
	e2eutil.RunCommand(t, "dart", "run", "./tests/dart/run.dart")
}
