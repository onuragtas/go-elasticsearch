package go_elasticsearch

type Operation struct {
	Adapter IAdapter
}

func (o Operation) UpdateByQuery(query Main) ([]byte, error) {
	return o.Adapter.UpdateByQuery(query)
}

func (o Operation) UpdateWithId(id string, source map[string]interface{}) ([]byte, error) {
	return o.Adapter.UpdateWithId(id, source)
}

func (o Operation) AddToExists(to []map[string]interface{}, value interface{}) []map[string]interface{} {
	return o.Adapter.AddToExists(to, value)
}

func (o Operation) AddToTerm(to []map[string]interface{}, key string, value interface{}) []map[string]interface{} {
	return o.Adapter.AddToTerm(to, key, value)
}

func (o Operation) AddToTerms(to []map[string]interface{}, key string, value interface{}) []map[string]interface{} {
	return o.Adapter.AddToTerms(to, key, value)
}

func (o Operation) AddToObject(to []map[string]interface{}, obj map[string]interface{}) []map[string]interface{} {
	return o.Adapter.AddToObject(to, obj)
}

func (o Operation) AddToRange(slice []map[string]interface{}, key string, from, to interface{}) []map[string]interface{} {
	return o.Adapter.AddToRange(slice, key, from, to)
}

func (o Operation) Scroll(main Main) (Result, error) {
	return o.Adapter.Scroll(main)
}

func (o Operation) ScrollById(result Result) (Result, error) {
	return o.Adapter.ScrollById(result)
}

func (o Operation) Search(main Main) (Result, error) {
	return o.Adapter.Search(main)
}

func (o Operation) Bulk(main interface{}) ([]byte, error) {
	return o.Adapter.Bulk(main)
}

func NewOperation(adapter IAdapter) IOperation {
	return Operation{Adapter: adapter}
}
