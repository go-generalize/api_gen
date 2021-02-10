// Package main ...
package main

// RequestParam ...
type RequestParam struct {
	Name     string
	Type     RequestParamType
	DataType string
	Comment  string
}

// RequestParamType ...
type RequestParamType string

// request param types
const (
	QueryRequestName RequestParamType = "query"
	PathRequestName  RequestParamType = "path"
	BodyRequestName  RequestParamType = "body"
)

// data types
const (
	StringRequestDataType  string = "string"  // string
	IntegerRequestDataType string = "integer" // int, uint, uint32, uint64
	NumberRequestDataType  string = "number"  // float32
	BooleanRequestDataType string = "boolean" // bool
)
