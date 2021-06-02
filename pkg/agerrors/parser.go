// Package agerrors has error types for api_gen
package agerrors

import "fmt"

// NewParserError initializes a new ParserError
func NewParserError(file string, line int, message string) error {
	return &ParserError{
		File:    file,
		Line:    line,
		Message: message,
	}
}

// ParserError represents an error occurred in parser
type ParserError struct {
	File    string
	Line    int
	Message string
}

var _ error = &ParserError{}
var _ Formatted = &ParserError{}

// Error is for error interface
func (e *ParserError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Message)
}

// Error is for Formatted interface
func (e *ParserError) Formatted() string {
	return e.Error()
}
