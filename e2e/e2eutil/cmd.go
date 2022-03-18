package e2eutil

import (
	"os/exec"
	"testing"
)

// RunCommand runs command with args.
func RunCommand(t *testing.T, name string, arg ...string) {
	t.Helper()
	cmd := exec.Command(name, arg...)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to run command %s %v: %v(%s)", name, arg, err, output)
	}
}

// RunCommandWithModifier runs command with modifier to customize exec.Cmd.
func RunCommandWithModifier(t *testing.T, modifier func(cmd *exec.Cmd), name string, arg ...string) {
	t.Helper()
	cmd := exec.Command(name, arg...)
	modifier(cmd)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to run command %s %v: %v(%s)", name, arg, err, output)
	}
}
