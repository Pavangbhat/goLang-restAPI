package models

type Goals struct {
	Id     int64 `json:"id"`
	Status string `json:"status"`
	Title string `json:"title"`
}