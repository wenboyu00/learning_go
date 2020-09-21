package parser

import (
	"bytes"
	"github.com/antchfx/htmlquery"
	"learning_go/crawler/engine"
)

func ParseCityList(body []byte) engine.ParseResult {
	r := bytes.NewReader(body)
	doc, err := htmlquery.Parse(r)
	if err != nil {
		panic(err)
	}
	list := htmlquery.Find(doc, `//*[@id="app"]/article[2]/dl/dd/a`)
	result := engine.ParseResult{}
	for _, n := range list {
		name := htmlquery.InnerText(n)
		url := htmlquery.SelectAttr(n, "href")
		result.Items = append(result.Items, string(name))
		result.Requests = append(result.Requests,
			engine.Request{Url: url, ParserFun: engine.NilParser})
	}
	return result
}
