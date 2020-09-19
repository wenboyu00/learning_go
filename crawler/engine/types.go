package engine

import (
	"io"
)

type ParseResult struct {
	Requests	[]Request
	Items	[]interface{}
}

type Request struct {
	Url 	string
	ParserFun	func(body io.ReadCloser) ParseResult
}

func NilParser(body io.ReadCloser) ParseResult  {
	return ParseResult{}
}