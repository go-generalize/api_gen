// Package main is the entrypoint for  api_gen cmd
package main

import (
	"fmt"

	"github.com/go-generalize/api_gen/v2/common"
	"github.com/spf13/cobra"
)

func init() {
	versionCommand := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Show version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("api_gen version: %s\n", common.AppVersion)
		},
	}

	rootCmd.AddCommand(versionCommand)
}
