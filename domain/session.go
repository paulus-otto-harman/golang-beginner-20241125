package domain

import "time"

type Session struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}
