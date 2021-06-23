package user

type RegisterInput struct {
	FullName string `json:"full_name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type InputLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
