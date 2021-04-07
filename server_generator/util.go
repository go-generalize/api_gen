// Package main ...
package main

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/go-utils/gopackages"
	"golang.org/x/xerrors"
)

// ExecCommand - コマンドを実行して出力結果とエラーを返す
func ExecCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	b, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	if exitCode := cmd.ProcessState.ExitCode(); exitCode != 0 {
		return "", fmt.Errorf("failed to exec git command: (exit code: %d, output: %s)", exitCode, string(b))
	}

	return string(b), nil
}

// GetGoRootPath - go.modがあるフォルダを返す
func GetGoRootPath(in string) (string, error) {
	d, err := gopackages.GetGoModPath(in)
	if err != nil {
		return "", xerrors.Errorf("failed to get go.mod directory: %w", err)
	}
	return filepath.Dir(d), nil
}

// GetGoRootPackageName - Goのルートパッケージ名をgo.modから取得する
func GetGoRootPackageName(in string) (string, error) {
	fp, err := gopackages.GetGoModPath(in)
	if err != nil {
		return "", xerrors.Errorf("failed to get go.mod file: %w", err)
	}

	p, err := gopackages.GetGoModule(fp)
	if err != nil {
		return "", xerrors.Errorf("failed to read go.mod: %w", err)
	}
	return p, nil
}
