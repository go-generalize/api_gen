// Package clientcmd has commands for "api_gen client"
package clientcmd

import (
	"os"
	"path/filepath"

	"github.com/go-generalize/api_gen/common"
	clientts "github.com/go-generalize/api_gen/pkg/client/typescript"
	"github.com/go-generalize/api_gen/pkg/parser"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

var tsCommand = func() *cobra.Command {
	var dir string

	cmd := &cobra.Command{
		Use:     "typescript [path]",
		Aliases: []string{"ts"},
		Short:   "Generate TypeScript client library",
		Long: `Generate TypeScript client library.
Pass the directory to parse as the 1st argument.`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			group, err := parser.Parse(args[0])

			if err != nil {
				return xerrors.Errorf("failed to parse the package(%s): %w", args[0], err)
			}

			generator := clientts.NewGenerator(group, common.AppVersion)

			code, err := generator.GenerateClient()

			if err != nil {
				return xerrors.Errorf("failed to generate source code: %w", err)
			}

			path := filepath.Join(dir, "api_client.ts")
			if err := os.WriteFile(path, []byte(code), 0664); err != nil {
				return xerrors.Errorf("failed to save in %s: %w", path, err)
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
	cmd.Flags().StringVarP(&dir, "file", "o", "./", "The directory to generate client library in")

	return cmd
}()
