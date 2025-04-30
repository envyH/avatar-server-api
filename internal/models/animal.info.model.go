package models

type AnimalInfo struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	BornTime   int    `json:"born_time"`
	MatureTime int    `json:"mature_time"`
}
