package repository

import "gin-person/model"

type IPersonRepository interface {
	Create(newPerson model.Person) (model.Person, error)
	GetAll() ([]model.Person, error)
	Update(person model.Person) (model.Person, error)
	Delete(id int) error
}