package adapters

import (
	"encoding/json"
	"github.com/onuragtas/go_elasticsearch"
)

func Decorate(res []byte) (go_elasticsearch.Result, error) {
	var result go_elasticsearch.Result

	err := json.Unmarshal(res, &result)

	return result, err
}
