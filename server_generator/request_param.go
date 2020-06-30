package main

type RequestParam struct {
	Name     string
	Type     RequestParamType
	DataType string
}

type RequestParamType string

const (
	QueryRequestName RequestParamType = "query"
	PathRequestName  RequestParamType = "path"
	BodyRequestName  RequestParamType = "body"
)

const (
	StringRequestDataType  string = "string"  //string
	IntegerRequestDataType string = "integer" // int, uint, uint32, uint64
	NumberRequestDataType  string = "number"  // float32
	BooleanRequestDataType string = "boolean" // bool
)
