package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"latihan-web-server/formatter"
	"latihan-web-server/helper"
	"latihan-web-server/usecase"
	"net/http"
)

type PersonHandler interface {
	AddPerson(w http.ResponseWriter, r *http.Request)
	GetPersons(w http.ResponseWriter, r *http.Request)
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

	var p usecase.InputPerson

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error: %s", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }
	
	json.Unmarshal(reqBody, &p)

	result, err := ph.person.AddPerson(p)
	if err != nil {
		fmt.Printf("error: %s", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

	pFormatter := formatter.ResponseFormatterPerson(result)
	dataResponse := helper.ApiResponse("Success Add Person", http.StatusOK, "success", pFormatter)
	json.NewEncoder(w).Encode(dataResponse)
}

func (ph *personHandler) GetPersons(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

	dataPerson, err := ph.person.GetPersons()
	if err != nil {
		fmt.Printf("error: %s", err)
        w.WriteHeader(http.StatusNotFound)
        return
	}

	json.NewEncoder(w).Encode(dataPerson)
}