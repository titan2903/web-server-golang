package usecase

import (
	"encoding/hex"
	"latihan-web-server/entity"
	"latihan-web-server/repository"
	"latihan-web-server/transport"

	uuid "github.com/satori/go.uuid"
)

type PersonService interface {
	AddPerson(data transport.InputPerson) (entity.Person, error)
	GetPersons() []entity.Person
	GetPersonById(id string) (entity.Person, error)
	// DeletePersonById(id string) (string, error)
}

type personService struct {
	person repository.PersonRepository
}

func NewPersonUsecase(pRepository repository.PersonRepository) PersonService {
	return &personService{person: pRepository}
}

func (ps *personService) AddPerson(data transport.InputPerson) (entity.Person, error) {
	id := uuid.NewV4()

	dataCreated := entity.Person{}
	dataCreated.ID = hex.EncodeToString(id[:])
	dataCreated.Name = data.Name
	dataCreated.Gender = data.Gender
	dataCreated.Height = data.Height
	dataCreated.Age = data.Age

	personDesc, err := ps.person.AddPerson(dataCreated)
	if err != nil {
		return personDesc, err
	}

	return personDesc, nil
}

func (ps *personService) GetPersons() []entity.Person {
	pList := ps.person.GetPersons()
	return pList
}

func (ps *personService) GetPersonById(id string) (entity.Person, error) {
	p, err := ps.person.GetPersonById(id)
	if err != nil {
		return p, err
	}

	return p, nil
}

// func (ps *personService) DeletePersonById(id stirng) (string, error) {
	
// }