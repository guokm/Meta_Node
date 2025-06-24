package dto

type RegUser struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,min=6"`
}

type UserLogin struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}
