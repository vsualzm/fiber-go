package request

type UserCreateRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"adress"`
	Phone   string `json:"phone"`
}
