package domain

import "time"

type WishlistItem struct {
	ProductId int       `json:"product_id" validate:"required,exists=products#id"`
	CreatedAt time.Time `json:"created_at"`
}
