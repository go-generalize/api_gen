// Package clientcmd has commands for "api_gen client"
package clientcmd

import (
	"os"
	"path/filepath"

	"github.com/go-utils/gopackages"

	"github.com/go-generalize/api_gen/common"
	clientgo "github.com/go-generalize/api_gen/pkg/client/go"
	"github.com/go-generalize/api_gen/pkg/parser"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

var goCommand = func() *cobra.Command {
	var dir, pkg string

	cmd := &cobra.Command{
		Use:   "go [path]",
		Short: "Generate go client library",
		Long: `Generate go client library.
Pass the directory to parse as the 1st argument.`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			group, err := parser.Parse(args[0])

			if err != nil {
				return xerrors.Errorf("failed to parse the package(%s): %w", args[0], err)
			}

			classesRoot := filepath.Join(dir, "classes")

			mod, err := gopackages.NewModule(classesRoot)

			if err != nil {
				return xerrors.Errorf("failed to analyze module in %s: %w", dir, err)
			}

			importPath, err := mod.GetImportPath(classesRoot)

			if err != nil {
				return xerrors.Errorf("failed to get import path for %s: %w", dir, err)
			}

			generator := clientgo.NewGenerator(group, importPath, pkg, common.AppVersion)

			code, err := generator.GenerateClient()

			if err != nil {
				return xerrors.Errorf("failed to generate source code: %w", err)
			}

			file := filepath.Join(dir, "api_client.go")
			if err := os.WriteFile(file, []byte(code), 0664); err != nil {
				return xerrors.Errorf("failed to save in %s: %w", file, err)
			}

			err = generator.GenerateTypes(func(relPath, code string) error {
				path := filepath.Join(dir, relPath)
				dir := filepath.Dir(path)

				if err := os.MkdirAll(dir, 0774); err != nil {
					return xerrors.Errorf("failed to mkdir %s recursively: %w", dir, err)
				}

				if err := os.WriteFile(path, []byte(code), 0664); err != nil {
					return xerrors.Errorf("failed to save code in %s: %w", path, err)
				}

				return nil
			})

			if err != nil {
				return xerrors.Errorf("failed to generate types: %w", err)
			}

			return nil
		},
	}
	cmd.Flags().StringVarP(&dir, "output", "o", "./", "The directory to generated client library in")
	cmd.Flags().StringVarP(&pkg, "package", "p", "client", "The package name of the generated library")

	return cmd
}()
