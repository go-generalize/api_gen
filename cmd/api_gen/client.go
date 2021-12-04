// Package main is the entrypoint for  api_gen cmd
package main

import (
	clientcmd "github.com/go-generalize/api_gen/v2/cmd/api_gen/client"
	"github.com/spf13/cobra"
)

func init() {
	clientCommand := &cobra.Command{
		Use:     "client",
		Aliases: []string{"c"},
		Short:   "Generates client-side library",
	}

	for _, cmd := range clientcmd.Initialize() {
		clientCommand.AddCommand(cmd)
	}

	rootCmd.AddCommand(clientCommand)
}
