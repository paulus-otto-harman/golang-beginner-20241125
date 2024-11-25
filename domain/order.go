package domain

type Order struct {
	CustomerToken string `json:"customer_token" validate:"not-empty-cart"`
	Id            int    `json:"id"`
	AddressId     int    `json:"address_id" validate:"required,exists-for-user=CustomerToken"`
	CouponCode    string `json:"coupon_code" validate:"max=20"`
	Shipping      string `json:"shipping" validate:"required,is-regular"`
	Notes         string `json:"notes" validate:"max=128"`
	PaymentMethod string `json:"payment_method" validate:"oneof=bank cod"`
	CreatedAt     string `json:"created_at"`
}
