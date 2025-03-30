package models

type Notification struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}
