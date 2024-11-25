package domain

type Recommendation struct {
	Title     *string `json:"title"`
	Subtitle  *string `json:"subtitle"`
	Photo     *string `json:"photo"`
	ProductId int     `json:"product_id"`
}
