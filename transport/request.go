package transport

type InputPerson struct{
	Name string `json:"name" validate:"required"`
	Gender string `json:"gender" validate:"required"`
	Height int64 `json:"height"`
	Age int64 `json:"age"`
}

type InputIdPerson struct{
	ID string `uri:"id"`
}