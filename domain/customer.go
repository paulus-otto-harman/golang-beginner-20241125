package domain

import "time"

type Customer struct {
	Name      string    `json:"name" validate:"required,max=30"`
	Username  string    `json:"username,omitempty" validate:"required,unique=users#username,max=40"`
	Password  string    `json:"password,omitempty" validate:"required,max=16"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
