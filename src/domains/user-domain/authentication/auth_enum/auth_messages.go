package auth_enum

type ErrorMessage string
const (
	InvalidCredentials ErrorMessage  = "Login creadentials are invalid"
	EmailExist ErrorMessage  = "Email already exists"
)
