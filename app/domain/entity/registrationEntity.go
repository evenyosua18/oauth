package entity

type RegistrationUserRequest struct {
	Name     string
	Password string
	Email    string
	Phone    string
}

type RegistrationUserResponse struct {
	Id string
}
