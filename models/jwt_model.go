package models

type JwtModel struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
