package adapters

import (
	"encoding/json"
	"go-elasticsearch"
)

func Decorate(res []byte) (go_elasticsearch.Result, error) {
	var result go_elasticsearch.Result

	err := json.Unmarshal(res, &result)

	return result, err
}
