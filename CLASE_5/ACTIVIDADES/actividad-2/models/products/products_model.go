package products

import "context"

type Query struct {
	Text     string
	Category string
	Brand    string
	Limit    int
}

type Result struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Brand    string  `json:"brand"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Score    float64 `json:"score"`
}

type Engine interface {
	Search(context.Context, Query) ([]Result, error)
}
