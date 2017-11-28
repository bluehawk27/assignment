package store

type Message struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
