package handlers

type UserRegister struct {
	Username string `validate:"required,min=5,max=20"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
type UserByEmail struct {
	Username string `validate:"required,min=5,max=20"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
	Type     string
	Image    string
	Token_id string
}

type UserLogin struct {
	Email    string `json:"email"    validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required"`
}

// Define a struct to hold the JSON data
type UserData struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}
