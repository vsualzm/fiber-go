package request

type UserCreateRequest struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UserUpdateRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type UserEmailRequest struct {
	Email string `json:"email" validate:"required"`
}
