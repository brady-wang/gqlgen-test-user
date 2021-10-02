package model

type User struct {
	ID     int `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Hobby  string `json:"hobby"`
}