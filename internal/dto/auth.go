package dto

type RegisterUser struct {
	FirstName string `json:"first_name" binding:"required,min=3"`
	LastName  string `json:"last_name" binding:"required,min=3"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}