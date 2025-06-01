package types

type Student struct{
	Id int 
	Name string
	Email string
	Age int

}

type ErrorResponse struct{
	Status string
	Error string
}