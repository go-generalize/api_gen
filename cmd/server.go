// Package main is the entrypoint for  api_gen cmd
package main

import (
	servercmd "github.com/go-generalize/api_gen/cmd/server"
	"github.com/spf13/cobra"
)

func init() {
	serverCommand := &cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "Generates server-side library",
	}

	for _, cmd := range servercmd.Initialize() {
		serverCommand.AddCommand(cmd)
	}

	rootCmd.AddCommand(serverCommand)
}
