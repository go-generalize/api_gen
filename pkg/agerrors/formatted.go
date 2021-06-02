// Package agerrors has error types for api_gen
package agerrors

import (
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

// Formatted is a formatted error representation which is printed directly
type Formatted interface {
	error
	Formatted() string
}

// UnwrapFormattedRunE wraps RunE and if returned error coontains Formatted, it will be returned
func UnwrapFormattedRunE(fn func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		err := fn(cmd, args)

		var fmtd Formatted
		if xerrors.As(err, &fmtd) {
			return fmtd
		}

		return err
	}
}
