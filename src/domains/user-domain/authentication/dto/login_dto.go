package authentication

type LoginDto struct {
	Identifier string `json:"identifier" validate:"required"` 
	Password string `json:"password" validate:"required,min=8,max=16"` 
}

type RegisterDto struct {
	Email string `json:"email" validate:"required,email"` 
	Username string `json:"username" validate:"required"` 
	Password string `json:"password" validate:"required,min=8,max=16"` 
}
