package Dto

type UserDTOInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserDTOOutput struct {
	Id       string
	Name     string
	Email    string
	Password string
}
