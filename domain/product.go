package domain

type Products []Products

type Product struct {
	ID       uint    `json:"id"`
	Name     string  `json:"product"`
	Category string  `json:"category"`
	ImageURL string  `json:"image"`
	Price    float64 `json:"price"`
}
