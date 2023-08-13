package types

type Sort struct {
	Text []string `json:"text"`
}

type Size struct {
	From int64 `json:"from"`
	Size int64 `json:"size"`
}
