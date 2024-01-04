package auth_dto

type LoginDto struct {
	Username string `json:"username" validate:"required,min=8,max=16"` 
	Password string `json:"password" validate:"required,min=8,max=16"` 
}