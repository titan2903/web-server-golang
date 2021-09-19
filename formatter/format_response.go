package formatter

import (
	"latihan-web-server/entity"
)

type Response struct {
	ID string`json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Height int64 `json:"height"`
	Age int64 `json:"age"`
}

func ResponseFormatterPerson(person entity.Person) Response {
	format := Response{}
	format.ID = person.ID
	format.Name = person.Name
	format.Age = person.Age
	format.Gender = person.Gender
	format.Height = person.Height

	return format
}
