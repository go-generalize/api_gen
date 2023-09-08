// Package main is the entrypoint for api_gen cmd
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "api_gen",
		Short: "Generates client/server API library from definitions in Go",
		Long: `api_gen generates client/server API library from definitions in Go.
GitHub: https://github.com/go-generalize/api_gen/v2`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
