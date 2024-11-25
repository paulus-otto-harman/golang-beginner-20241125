package domain

type Weekly struct {
	Photo     *string `json:"photo"`
	Title     *string `json:"title"`
	Subtitle  *string `json:"subtitle"`
	ProductId int     `json:"product_id"`
}
