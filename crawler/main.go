package main

import (
	"learning_go/crawler/engine"
	"learning_go/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParserFun: parser.ParseCityList,
	})
}