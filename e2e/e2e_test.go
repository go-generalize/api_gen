package e2etest

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/go-gneralize/api_gen/v2/e2e/e2eutil"
)

var apiGenCommand = func() string {
	if cmd := os.Getenv("E2E_API_GEN_COMMAND"); cmd != "" {
		return cmd
	}

	return "../bin/api_gen"
}()

var (
	wd = func() string {
		wd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		return wd
	}()
)

func runAPIGenServer(t *testing.T, target string) {
	t.Helper()

	b, err := exec.Command(
		apiGenCommand,
		"server",
		"-o",
		filepath.Join(target, "controller"),
		filepath.Join(target, "api"),
	).CombinedOutput()

	if err != nil {
		t.Fatal(err, string(b))
	}
}

func runAPIGenClient(t *testing.T, target string, clients ...string) {
	t.Helper()

	os.RemoveAll(filepath.Join(target, "clients"))

	for _, c := range clients {
		b, err := exec.Command(
			apiGenCommand,
			"client",
			c,
			"-o",
			filepath.Join(target, "clients", c),
			filepath.Join(target, "api"),
		).CombinedOutput()

		if err != nil {
			t.Fatal(c, err, string(b))
		}
	}
}

func TestMultipart(t *testing.T) {
	t.Helper()

	runAPIGenServer(t, "multipart")
	runAPIGenClient(t, "multipart", "ts", "dart", "go")

	if os.Getenv("E2E_API_GENERATE_ONLY") == "1" {
		return
	}

	e2eutil.RunCommandWithModifier(t, func(cmd *exec.Cmd) {
		cmd.Env = append(os.Environ(), "E2E_API_GEN_ROOT_DIR="+wd)
	}, "go", "test", "-v", "-tags=internal", "./multipart/...")
}
