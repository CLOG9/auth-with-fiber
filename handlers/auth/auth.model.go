package handlers

type UserRegister struct {
	Username string `validate:"required,min=5,max=20"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
type UserLogin struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required"`
}
