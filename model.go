package go_elasticsearch

type IOperation interface {
	Bulk(main interface{}) ([]byte, error)
	Search(main Main) (Result, error)
	Scroll(main Main) (Result, error)
	ScrollById(result Result) (Result, error)
	AddToTerm(to []map[string]interface{}, key string, value interface{}) []map[string]interface{}
	AddToTerms(to []map[string]interface{}, key string, value interface{}) []map[string]interface{}
	AddToExists(to []map[string]interface{}, value interface{}) []map[string]interface{}
	AddToRange(slice []map[string]interface{}, key string, from, to interface{}) []map[string]interface{}
	UpdateWithId(id string, source map[string]interface{}) ([]byte, error)
	UpdateByQuery(query Main) ([]byte, error)
	AddToObject(slice []map[string]interface{}, obj map[string]interface{}) []map[string]interface{}
}

type IAdapter interface {
	Bulk(main interface{}) ([]byte, error)
	Search(main Main) (Result, error)
	Scroll(main Main) (Result, error)
	ScrollById(result Result) (Result, error)
	AddToTerm(to []map[string]interface{}, key string, value interface{}) []map[string]interface{}
	AddToTerms(to []map[string]interface{}, key string, value interface{}) []map[string]interface{}
	AddToExists(to []map[string]interface{}, value interface{}) []map[string]interface{}
	AddToRange(slice []map[string]interface{}, key string, from, to interface{}) []map[string]interface{}
	UpdateWithId(id string, source map[string]interface{}) ([]byte, error)
	UpdateByQuery(query Main) ([]byte, error)
	AddToObject(slice []map[string]interface{}, obj map[string]interface{}) []map[string]interface{}
}

type Main struct {
	Size         int                      `json:"size,omitempty"`
	From         int                      `json:"from,omitempty"`
	Query        Query                    `json:"query"`
	Aggregations map[string]interface{}   `json:"aggs,omitempty"`
	Script       map[string]string        `json:"script,omitempty"`
	Sort         []map[string]interface{} `json:"sort,omitempty"`
}

type D map[string]interface{}

type Query struct {
	Bool Bool `json:"bool"`
}

type Bool struct {
	Must    Must                   `json:"must"`
	MustNot MustNot                `json:"must_not"`
	Should  Should                 `json:"should"`
	Filter  map[string]interface{} `json:"filter"`
}

type Must []map[string]interface{}

type MustNot []map[string]interface{}

type Should []map[string]interface{}

type Term struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type Range struct {
	Key  string      `json:"key"`
	From interface{} `json:"from"`
	To   interface{} `json:"to"`
}

type Result struct {
	ScrollID string `json:"_scroll_id"`
	Took     int    `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string                 `json:"_index"`
			Type   string                 `json:"_type"`
			ID     string                 `json:"_id"`
			Score  float64                `json:"_score"`
			Source map[string]interface{} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
	Aggregations map[string]interface{} `json:"aggregations"`
}
