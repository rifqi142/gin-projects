package model

var Persons []Person = make([]Person, 0)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}