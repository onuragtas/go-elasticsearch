package go_elasticsearch

type IOperation interface {
	Search(main Main) (Result, error)
	Scroll(main Main) (Result, error)
	ScrollById(result Result) (Result, error)
	AddToTerm(to []map[string]interface{}, key string, value interface{}) []map[string]interface{}
	AddToExists(to []map[string]interface{}, value interface{}) []map[string]interface{}
	AddToRange(slice []map[string]interface{}, key string, from, to interface{}) []map[string]interface{}
	UpdateWithId(id string, source map[string]interface{}) ([]byte, error)
	UpdateByQuery(query Main) ([]byte, error)
}

type Operation struct {
	Adapter IOperation
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

func NewOperation(adapter IOperation) IOperation {
	return Operation{Adapter: adapter}
}
