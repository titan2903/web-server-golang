package usecase

import (
	"latihan-web-server/entity"
	"latihan-web-server/repository"
)

type InputPerson struct{
	Name string `json:"name"`
	Gender string `json:"gender"`
	Height int64 `json:"height"`
	Age int64 `json:"age"`
}

type PersonService interface {
	AddPerson(data InputPerson) (entity.Person, error)
	GetPersons() ([]entity.Person, error)
}

type personService struct {
	person repository.PersonRepository
}

func NewPersonUsecase(pRepository repository.PersonRepository) PersonService {
	return &personService{person: pRepository}
}

func (ps *personService) AddPerson(data InputPerson) (entity.Person, error) {
	dataCreated := entity.Person{}
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

func (ps *personService) GetPersons() ([]entity.Person, error) {
	p, err := ps.person.GetPersons()
	if err != nil {
		return p, err
	}

	return p, nil
}