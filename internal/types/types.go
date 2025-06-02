package types

type Student struct {
	Id    int
	Name  string `validate:required`
	Email string `validate:required`
	Age   int    `validate:required`
}

type ErrorResponse struct {
	Status string
	Error  string
}
