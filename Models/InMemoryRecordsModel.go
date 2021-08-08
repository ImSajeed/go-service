package models

type InMemoryRecordsRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
