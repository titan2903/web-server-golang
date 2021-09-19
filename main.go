package main

import (
	"latihan-web-server/handler"
	"latihan-web-server/repository"
	"latihan-web-server/usecase"
	"log"
	"net/http"
)

func main()  {
	
	pRepository := repository.NewPersonRepository()
	pUseCase := usecase.NewPersonUsecase(pRepository)
	pHandler := handler.NewPersonHandler(pUseCase)
	
	http.HandleFunc("/add", pHandler.AddPerson)
	http.HandleFunc("/persons", pHandler.GetPersons)
	http.HandleFunc("/person/:id", pHandler.GetPersonById)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}