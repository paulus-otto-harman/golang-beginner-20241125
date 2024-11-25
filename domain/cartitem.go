package domain

type CartItem struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
