package domain

type Cart struct {
	Items     []CartItem `json:"items"`
	ItemCount int        `json:"item_count"`
}
