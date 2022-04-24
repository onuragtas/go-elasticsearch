#Install
```
go get github.com/onuragtas/go_elasticsearch
```

#Usage
```
import (
	"github.com/onuragtas/go_elasticsearch"
	v7 "github.com/onuragtas/go_elasticsearch/adapters/v7"
)
```

```
func TestSearchV7(t *testing.T) {
	adapter := v7.NewElasticSearch("http://localhost:9200", "epa", "_doc", 0, 100)
	operation := go_elasticsearch.NewOperation(adapter)

	mainQuery := go_elasticsearch.Main{}
	mainQuery.Query.Bool.Must = operation.AddToTerm(mainQuery.Query.Bool.Must, "id", 99516062)

	result, err := operation.Search(mainQuery)
	log.Println(result, err)
}
```