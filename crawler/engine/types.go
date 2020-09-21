package engine

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

type Request struct {
	Url       string
	ParserFun func(body []byte) ParseResult
}

func NilParser(body []byte) ParseResult {
	return ParseResult{}
}
