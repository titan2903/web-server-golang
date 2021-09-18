package repository

import (
	"errors"
	"fmt"
	"latihan-web-server/entity"
)

type PersonRepository interface {
	AddPerson(data entity.Person) (entity.Person, error)
	GetPersons() ([]entity.Person, error)
}

type personRepository struct {
	person []entity.Person
}

func NewPersonRepository() PersonRepository {
	return &personRepository{}
}

func(pr *personRepository) AddPerson(data entity.Person) (entity.Person, error) {
	err := append(pr.person, data)
	if err == nil {
		fmt.Printf("Error: %+v", err)
		return data, errors.New("error add person in repository")
	}

	return data, nil
}

func(pr *personRepository) GetPersons() ([]entity.Person, error) {
	persons := []entity.Person{
		{Name: "john", Gender: "male", Height: 168, Age: 23},
		{Name: "jane", Gender: "female", Height: 170, Age: 22},
		{Name: "alex", Gender: "male", Height: 180, Age: 30},
	}

	return persons, nil
}