package domain

type Address struct {
	Id        int    `json:"id"`
	FullName  string `json:"full_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Detail    string `json:"detail" validate:"required"`
	IsDefault bool   `json:"is_default"`
}
