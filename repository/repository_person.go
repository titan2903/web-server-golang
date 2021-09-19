package repository

import (
	"errors"
	"latihan-web-server/entity"
)

type PersonRepository interface {
	AddPerson(data entity.Person) (entity.Person, error)
	GetPersons() []entity.Person
	GetPersonById(id string) (entity.Person, error)
	// DeletePersonById(id string) (string, error)
}

type personRepository struct {
	person []entity.Person
}

func NewPersonRepository() PersonRepository {
	return &personRepository{}
}

func(pr *personRepository) AddPerson(data entity.Person) (entity.Person, error) {
	pr.person = append(pr.person, data)
	return data, nil
}

func(pr *personRepository) GetPersons() []entity.Person{
	return pr.person
}

func(pr *personRepository) GetPersonById(id string) (entity.Person, error) {

	var dataPerson entity.Person

	var persons []entity.Person

	for _, person := range persons {
		if person.ID == id {
			dataPerson = person
		} else {
			return person, errors.New("Data Not Found")
		}

	}

	return dataPerson, nil
}

// func(pr *personRepository) DeletePersonById(id string) (string, error) {

// 	for i, person := range persons {
// 		if person.ID == id {
// 			persons = append(persons[:i], persons[i+1:]...)
// 			return "success deleted data" , nil
// 		}
// 	}

// 	return "", nil
// }