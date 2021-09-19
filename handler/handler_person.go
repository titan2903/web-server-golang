package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"latihan-web-server/formatter"
	"latihan-web-server/transport"
	"latihan-web-server/usecase"
	"latihan-web-server/validation"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type PersonHandler interface {
	AddPerson(w http.ResponseWriter, r *http.Request)
	GetPersons(w http.ResponseWriter, r *http.Request)
	GetPersonById( w http.ResponseWriter, r *http.Request)
	// DeletePersonById( w http.ResponseWriter, r *http.Request)
}

type personHandler struct {
	person usecase.PersonService
}

func NewPersonHandler(pUseCase usecase.PersonService) PersonHandler {
	return &personHandler{person: pUseCase}
}

func (ph *personHandler) AddPerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

	w.Header().Set("Content-Type", "application/json")

	var p transport.InputPerson

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error: %s", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }
	
	json.Unmarshal(reqBody, &p)

	err = validator.New().Struct(p)
	if err != nil {
		errors := validation.FormatValidationError(err)
		fmt.Printf("error: %+v", errors)
		dataResponse := transport.ApiResponse(errors, http.StatusBadRequest, "err", nil)
		json.NewEncoder(w).Encode(dataResponse)
		return
	}

	result, err := ph.person.AddPerson(p)
	if err != nil {
		fmt.Printf("error: %s", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

	pFormatter := formatter.ResponseFormatterPerson(result)
	dataResponse := transport.ApiResponse("Success Add Person", http.StatusAccepted, "success", pFormatter)
	json.NewEncoder(w).Encode(dataResponse)
}

func (ph *personHandler) GetPersons(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

	w.Header().Set("Content-Type", "application/json")

	dataPersons := ph.person.GetPersons()

	dataResponse := transport.ApiResponse("Success Get Data Persons", http.StatusAccepted, "success", dataPersons)

	json.NewEncoder(w).Encode(dataResponse)
}

func (ph *personHandler) GetPersonById( w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/person/")
	fmt.Printf("id: %+v", id)

	dataPerson, err := ph.person.GetPersonById(id)
	if err != nil {
		fmt.Printf("error: %s", err)
        w.WriteHeader(http.StatusNotFound)
        return
	}

	json.NewEncoder(w).Encode(dataPerson)
}

func (ph *personHandler) DeletePersonById( w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

	id := r.URL.Query().Get("id")
	fmt.Printf("id: %+v", id)
}