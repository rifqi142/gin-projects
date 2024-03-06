package repository

import (
	"database/sql"
	"fmt"
	"gin-person/model"
)

type personRepository struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) *personRepository {
	return &personRepository{
		db: db,
	}
}

func (pr *personRepository) Create(newPerson model.Person) (model.Person, error) {
	query := "insert into person(name, address) values($1, $2) returning *"

	row := pr.db.QueryRow(query, newPerson.Name, newPerson.Address)
	err := row.Scan(&newPerson.Id, &newPerson.Name, &newPerson.Address)
	return newPerson, err
}

func (pr *personRepository) GetAll() ([]model.Person, error) {
	var persons = []model.Person{}

	query := "select * from person"
	rows, err := pr.db.Query(query)
	if err != nil {
		return persons, err
	}

	for rows.Next() {
		var p model.Person
		err := rows.Scan(&p.Id, &p.Name, &p.Address)
		if err != nil {
			fmt.Println(err)
			continue
		}

		persons = append(persons, p)
	}

	return persons, nil
}