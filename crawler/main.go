package main

import (
	"learning_go/crawler/engine"
	"learning_go/crawler/scheduler"
	"learning_go/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParserFun: parser.ParseCityList,
	})
}
