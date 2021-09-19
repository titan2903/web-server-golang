package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)



func FormatValidationError(err error) string {
	var errors []string
	var resultError string

	for _, e := range err.(validator.ValidationErrors) {
		// if e[1] ==
		errors = append(errors, e.Error())
	}

	var join string
	for _, str := range errors {
	  join += str
	}
  
	splited := strings.Split(join, ".")
	
	if splited[1] == "Gender' Error:Field validation for 'Gender' failed on the 'required' tag" {
		resultError = "Gender must be input"
	} else if splited[1] == "Name' Error:Field validation for 'Name' failed on the 'required' tag" {
		resultError = "Name must be input"
	} else {
		resultError = "Name and Gender must be input"
	}

	return resultError
}