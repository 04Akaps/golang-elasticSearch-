package schema

type User struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int64  `json:"age"`
}
