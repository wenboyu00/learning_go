package parser

import (
	"github.com/antchfx/htmlquery"
	"io"
	"learning_go/crawler/engine"
)

func ParseCityList(body io.ReadCloser) engine.ParseResult {
	doc, err := htmlquery.Parse(body)
	defer body.Close()
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
