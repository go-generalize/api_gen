package clientcmd

import (
	"os"

	"github.com/go-generalize/api_gen/common"
	clientgo "github.com/go-generalize/api_gen/pkg/client/go"
	"github.com/go-generalize/api_gen/pkg/parser"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

var tsCommand = func() *cobra.Command {
	var file, pkg string

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

			code, err := clientgo.Generate(group, pkg, common.AppVersion)

			if err != nil {
				return xerrors.Errorf("failed to generate source code: %w", err)
			}

			if err := os.WriteFile(file, []byte(code), 0664); err != nil {
				return xerrors.Errorf("failed to save in %s: %w", file, err)
			}

			return nil
		},
	}
	cmd.Flags().StringVarP(&file, "file", "o", "./client.go", "The path to generate client library in")

	return cmd
}()
