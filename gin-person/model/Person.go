package model

type Person struct {
	Id      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}