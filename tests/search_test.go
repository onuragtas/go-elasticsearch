package tests

import (
	"github.com/onuragtas/go_elasticsearch"
	v7 "github.com/onuragtas/go_elasticsearch/adapters/v7"
	"log"
	"testing"
)

func TestSearchV7(t *testing.T) {
	adapter := v7.NewElasticSearch("http://localhost:9200", "epa", "_doc", 0, 100)
	operation := go_elasticsearch.NewOperation(adapter)

	mainQuery := go_elasticsearch.Main{}
	mainQuery.Query.Bool.Must = operation.AddToTerm(mainQuery.Query.Bool.Must, "id", 99516062)

	sort := make(map[string]interface{})
	s := make(map[string]string)

	s["order"] = "asc"
	sort["id"] = s
	mainQuery.Sort = append(mainQuery.Sort, sort)

	result, err := operation.Search(mainQuery)
	log.Println(result, err)
}

func TestBulkV7(t *testing.T) {
	adapter := v7.NewElasticSearch("http://192.168.36.240:9200", "epa", "_doc", 0, 100)
	operation := go_elasticsearch.NewOperation(adapter)

	var bulk []interface{}
	bulk = append(bulk, go_elasticsearch.D{"index": go_elasticsearch.D{"_index": "epa", "_id": "test"}})
	bulk = append(bulk, go_elasticsearch.D{"field": "test"})

	bulk = append(bulk, go_elasticsearch.D{"index": go_elasticsearch.D{"_index": "epa", "_id": "test2"}})
	bulk = append(bulk, go_elasticsearch.D{"field": "test2"})

	bytes, err := operation.Bulk(bulk)
	if err != nil {
		return
	}

	log.Println(string(bytes), err)
}
