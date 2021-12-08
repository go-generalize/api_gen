// Package clientcmd has commands for "api_gen client"
package clientcmd

import "github.com/spf13/cobra"

// Initialize returns initialized api_gen client subcommands
func Initialize() []*cobra.Command {
	return []*cobra.Command{
		goCommand,
		tsCommand,
		dartCommand,
	}
}
