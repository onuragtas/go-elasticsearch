package go_elasticsearch

type IOperation interface {
	Search(main Main) (Result, error)
	Scroll(main Main) (Result, error)
	ScrollById(result Result) (Result, error)
	AddToTerm(to []map[string]interface{}, key string, value interface{}) []map[string]interface{}
	AddToRange(slice []map[string]interface{}, key string, from, to interface{}) []map[string]interface{}
}

type Operation struct {
	Adapter IOperation
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
