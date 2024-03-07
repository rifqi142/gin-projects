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

// create new data person
func (pr *personRepository) Create(newPerson model.Person) (model.Person, error) {
	query := "insert into person(name, address) values($1, $2) returning *"

	row := pr.db.QueryRow(query, newPerson.Name, newPerson.Address)
	err := row.Scan(&newPerson.Id, &newPerson.Name, &newPerson.Address)
	return newPerson, err
}

// get all data person
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

// update data person
func (pr *personRepository) Update(person model.Person) (model.Person, error) {
	query := "update person set name=$2, address=$3 where id=$1 returning *"

	row := pr.db.QueryRow(query, person.Id, person.Name, person.Address)
	var updatedPerson model.Person
	err := row.Scan(&updatedPerson.Id, &updatedPerson.Name, &updatedPerson.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Person{}, fmt.Errorf("person with ID %d not found", person.Id)
		}
		return model.Person{}, err
	}
	return updatedPerson, nil
}

// delete data person
func (pr *personRepository) Delete(id int) error {
	query := "delete from person where id=$1"

	_, err := pr.db.Exec(query, id)
	return err
}