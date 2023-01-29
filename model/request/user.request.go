package request

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Address  string `json:"address" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=32"`
	Phone    string `json:"phone" validate:"required"`
}

type UserUpdaterequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
