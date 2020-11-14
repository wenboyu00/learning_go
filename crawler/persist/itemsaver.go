package main

import (
	"github.com/elastic/go-elasticsearch/v6"
	"log"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200/",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	log.Println(elasticsearch.Version)
	log.Println(es.Info())
}
