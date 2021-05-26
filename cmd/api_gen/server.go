// Package main is the entrypoint for  api_gen cmd
package main

import (
	"github.com/go-generalize/api_gen/common"
	"github.com/go-generalize/api_gen/pkg/parser"
	"github.com/go-generalize/api_gen/pkg/server"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

func init() {
	var dir, basePackage string

	serverCommand := &cobra.Command{
		Use:     "server [path]",
		Aliases: []string{"s"},
		Short:   "Generates server-side library",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			group, err := parser.Parse(args[0])

			if err != nil {
				return xerrors.Errorf("failed to parse the package(%s): %w", args[0], err)
			}

			generator, err := server.NewGenerator(group, dir, basePackage, common.AppVersion)

			if err != nil {
				return xerrors.Errorf("failed to initialize generator: %w", err)
			}

			if err := generator.Generate(); err != nil {
				return xerrors.Errorf("failed to generate library: %w", err)
			}

			return nil
		},
	}
	serverCommand.Flags().StringVarP(&dir, "output", "o", "./", "The directory to generated server library in")
	serverCommand.Flags().StringVarP(&basePackage, "package", "p", "controller", "Package name of the base directory")

	rootCmd.AddCommand(serverCommand)
}
