package go_elasticsearch

type Main struct {
	Size   int                      `json:"size,omitempty"`
	From   int                      `json:"from,omitempty"`
	Query  Query                    `json:"query"`
	Script map[string]string        `json:"script,omitempty"`
	Sort   []map[string]interface{} `json:"sort,omitempty"`
}

type Query struct {
	Bool Bool `json:"bool"`
}

type Bool struct {
	Must    Must    `json:"must"`
	MustNot MustNot `json:"must_not"`
	Should  Should  `json:"should"`
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
}
