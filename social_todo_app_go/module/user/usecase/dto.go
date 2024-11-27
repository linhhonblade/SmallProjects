package usecase

type EmailPasswordRegistrationDTO struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type EmailPasswordLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
