// Package main ...
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

// GetGoRootPackageName - Goのルートパッケージ名をgo.modから取得する
func GetGoRootPackageName() (string, error) {
	goModPath := GetGoModPath()
	d, err := ioutil.ReadFile(goModPath)
	if err != nil {
		return "", err
	}

	f, err := modfile.Parse("", d, nil)
	if err != nil {
		return "", err
	}

	if len(f.Module.Mod.Path) == 0 {
		return "", fmt.Errorf("package name was not found")
	}

	return f.Module.Mod.Path, nil
}

// GetGoModPath - Gitリポジトリ直下から上がりながらgo.modを探す
func GetGoModPath() string {
	p := "./"

	for i := 0; i < 100; i++ {
		_, err := os.Stat(filepath.Join(p, "go.mod"))

		if err == nil {
			break
		}
		p = filepath.Join(p, "..")
	}

	p = filepath.Join(p, "go.mod")

	_, err := os.Stat(p)

	if err != nil {
		return ""
	}

	return p
}

// GetPackage returns package path to specified working directory
func GetPackage(p string) (string, error) {
	goMod := GetGoModPath()

	if goMod == "" {
		return "", fmt.Errorf("go.mod is not found in your repository")
	}

	packageName, err := GetGoRootPackageName()

	if err != nil {
		return "", fmt.Errorf("failed to get go root package name: %+v", err)
	}

	absGoMod, err := filepath.Abs(filepath.Dir(goMod))

	if err != nil {
		return "", fmt.Errorf("failed to get absolute path for go package root: %+v", err)
	}

	absCWD, err := filepath.Abs(p)

	if err != nil {
		return "", fmt.Errorf("failed to get absolute path for current directory: %+v", err)
	}

	rel, err := filepath.Rel(absGoMod, absCWD)

	if err != nil {
		return "", fmt.Errorf("failed to calc relative path from package root to here: %+v", err)
	}

	return path.Join(packageName, rel), nil
}
