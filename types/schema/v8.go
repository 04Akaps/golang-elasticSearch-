package schema

type V8ResponseType struct {
	Took    int64 `json:"took"`
	TimeOut bool  `json:"time_out"`
	Hits    Hits  `json:"hits"`
}

type Hits struct {
	Total    Total       `json:"total"`
	MaxScore float64     `json:"max_score"`
	Hits     []HitsTypes `json:"hits"`
}

type Total struct {
	Value    int64  `json:"value"`
	Relation string `json:"relation"`
}

type HitsTypes struct {
	Index  string      `json:"_index"`
	Type   string      `json:"_type"`
	Id     string      `json:"_id"`
	Score  float64     `json:"_score"`
	Source interface{} `json:"_source"`
}
