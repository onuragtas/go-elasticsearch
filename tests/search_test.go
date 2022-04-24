package tests

import (
	go_elasticsearch "go-elasticsearch"
	v7 "go-elasticsearch/adapters/v7"
	"log"
	"testing"
)

func TestSearchV7(t *testing.T) {
	adapter := v7.NewElasticSearch("http://localhost:9200", "epa", "_doc", 0, 100)
	operation := go_elasticsearch.NewOperation(adapter)

	mainQuery := go_elasticsearch.Main{}
	mainQuery.Query.Bool.Must = operation.AddToTerm(mainQuery.Query.Bool.Must, "id", 99516062)

	result, err := operation.Search(mainQuery)
	log.Println(result, err)
}
