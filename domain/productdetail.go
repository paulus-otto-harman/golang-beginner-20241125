package domain

type ProductDetail struct {
	Id           int              `json:"id"`
	Name         string           `json:"name"`
	ThumbnailUrl string           `json:"thumbnail_url"`
	Price        string           `json:"price"`
	DiscountRate *float64         `json:"discount_rate"`
	NetPrice     string           `json:"net_price"`
	CategoryId   int              `json:"category_id"`
	Rating       *int             `json:"rating"`
	OnMyWishlist *bool            `json:"on_my_wishlist"`
	CreatedAt    string           `json:"created_at"`
	IsNew        bool             `json:"is_new"`
	Variants     []ProductVariant `json:"variants"`
	Photos       []ProductPhoto   `json:"photos"`
}
