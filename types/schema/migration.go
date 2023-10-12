package schema

type Migration struct {
	PK      string `json:"pk"`
	Owner   string `json:"owner"`
	Tid     int64  `json:"tid"`
	KeyWord string `json:"keyWord"`
}
